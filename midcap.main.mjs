import { Client, GatewayIntentBits, EmbedBuilder, ActionRowBuilder, ButtonBuilder, ButtonStyle } from 'discord.js';
import { exec } from 'child_process';
import { promisify } from 'util';
import dotenv from 'dotenv';

dotenv.config();
const execPromise = promisify(exec);

const { DISCORD_TOKEN, CHANNEL_ID, SERVER_LOCATION } = process.env;

// Validate environment variables
if (!DISCORD_TOKEN || !CHANNEL_ID || !SERVER_LOCATION) {
  throw new Error('Missing required environment variables');
}

// Unique identifiers based on server location
const locationPrefix = SERVER_LOCATION.toLowerCase().replace(/\s+/g, '-');
const START_BUTTON_ID = `start-${locationPrefix}`;
const STOP_BUTTON_ID = `stop-${locationPrefix}`;

// Discord client setup
const client = new Client({
  intents: [
    GatewayIntentBits.Guilds,
    GatewayIntentBits.GuildMessages,
    GatewayIntentBits.GuildMessageReactions
  ]
});

// Check Docker status
async function isDockerRunning() {
  try {
    const { stdout } = await execPromise('docker ps -q -f name=hll-geofences-basic');
    return stdout.trim().length > 0;
  } catch (error) {
    console.error(`Error checking Docker status: ${error.message}`);
    return false;
  }
}

// Create embed
async function createEmbed() {
  const isRunning = await isDockerRunning();
  return new EmbedBuilder()
    .setTitle('Basic Seeding')
    .setDescription('Midcap only')
    .addFields({
      name: 'Docker Status',
      value: isRunning ? '🟢 Running' : '🔴 Stopped'
    })
    .setColor(isRunning ? 0x00FF00 : 0xFF0000)
    .setFooter({ text: `Server: ${SERVER_LOCATION}` })
    .setTimestamp();
}

// Create buttons
function createButtons() {
  return new ActionRowBuilder()
    .addComponents(
      new ButtonBuilder()
        .setCustomId(START_BUTTON_ID)
        .setLabel('START')
        .setStyle(ButtonStyle.Success),
      new ButtonBuilder()
        .setCustomId(STOP_BUTTON_ID)
        .setLabel('STOP')
        .setStyle(ButtonStyle.Danger)
    );
}

// Clear channel
async function clearChannel(channel) {
  try {
    const messages = await channel.messages.fetch({ limit: 100 });
    if (messages.size > 0) {
      await channel.bulkDelete(messages);
    }
  } catch (error) {
    console.error(`Error clearing channel: ${error.message}`);
  }
}

// Update embed
async function updateEmbed(channel) {
  try {
    const embed = await createEmbed();
    const buttons = createButtons();
    const messages = await channel.messages.fetch({ limit: 1 });
    const message = messages.first();

    if (message) {
      await message.edit({ embeds: [embed], components: [buttons] });
    } else {
      await channel.send({ embeds: [embed], components: [buttons] });
    }
  } catch (error) {
    console.error(`Error updating embed: ${error.message}`);
  }
}

// Execute docker command
async function executeDockerCommand(command, interaction) {
  try {
    // Only defer if the interaction has not been acknowledged
    if (!interaction.deferred && !interaction.replied) {
      await interaction.deferReply({ ephemeral: true });
    }

    // Run the Docker command
    const { stdout, stderr } = await execPromise(command);
    const output = stdout || stderr || "No output";

    // Only edit if the interaction was properly deferred
    await interaction.editReply({
      content: `Command executed successfully:\n\`\`\`\n${output.trim()}\n\`\`\``
    });

    // Update the channel embed
    await updateEmbed(interaction.channel);

  } catch (error) {
    console.error(`Error executing command: ${error.message}`);

    // If the interaction wasn't properly deferred, we need to catch this
    if (!interaction.deferred && !interaction.replied) {
      try {
        await interaction.deferReply({ ephemeral: true });
      } catch (ackError) {
        console.error(`Failed to defer interaction: ${ackError.message}`);
      }
    }

    // Edit the reply with the error message
    try {
      await interaction.editReply({
        content: `Error executing command: ${error.message}`
      });
    } catch (editError) {
      console.error(`Failed to edit reply: ${editError.message}`);
    }
  }
}

client.once('ready', async () => {
  console.log(`Bot started for ${SERVER_LOCATION}`);
  const channel = await client.channels.fetch(CHANNEL_ID);

  await clearChannel(channel);
  await updateEmbed(channel);

  setInterval(async () => {
    await updateEmbed(channel);
  }, 30000);
});

client.on('interactionCreate', async (interaction) => {
  if (!interaction.isButton()) return;

  if (interaction.customId === START_BUTTON_ID) {
    await executeDockerCommand('docker-compose up -d', interaction);
  } else if (interaction.customId === STOP_BUTTON_ID) {
    await executeDockerCommand('docker-compose down', interaction);
  }
});

client.login(DISCORD_TOKEN);

process.on('SIGTERM', async () => {
  console.log('Shutting down...');
  await client.destroy();
  process.exit(0);
});

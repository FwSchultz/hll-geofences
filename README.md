<div align="center">

  <img src="https://github.com/FwSchultz/assets/blob/main/bots/2KU77B0N3S/Logo.png" alt="logo" width="200" height="auto" />
  <h1>HLL-GEOFENCE</h1>
  
<!-- Badges -->
<p>
  <a href="https://github.com/2KU77B0N3S/hll-geofences/graphs/contributors">
    <img src="https://img.shields.io/github/contributors/2KU77B0N3S/hll-geofences" alt="contributors" />
  </a>
  <a href="">
    <img src="https://img.shields.io/github/last-commit/2KU77B0N3S/hll-geofences" alt="last update" />
  </a>
  <a href="https://github.com/2KU77B0N3S/hll-geofences/network/members">
    <img src="https://img.shields.io/github/forks/2KU77B0N3S/hll-geofences" alt="forks" />
  </a>
  <a href="https://github.com/2KU77B0N3S/hll-geofences/stargazers">
    <img src="https://img.shields.io/github/stars/2KU77B0N3S/hll-geofences" alt="stars" />
  </a>
  <a href="https://github.com/2KU77B0N3S/hll-geofences/issues/">
    <img src="https://img.shields.io/github/issues/2KU77B0N3S/hll-geofences" alt="open issues" />
  </a>
  <a href="https://github.com/2KU77B0N3S/hll-geofences/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/2KU77B0N3S/hll-geofences.svg" alt="license" />
  </a>
</p>
   
<h4>
  <a href="https://github.com/2KU77B0N3S/hll-geofences">Documentation</a>
  <span> · </span>
  <a href="https://github.com/2KU77B0N3S/hll-geofences/issues/">Report Bug</a>
  <span> · </span>
  <a href="https://github.com/2KU77B0N3S/hll-geofences/issues/">Request Feature</a>
</h4>
</div>

<br />

# Table of Contents

- [About the Project](#about-the-project)
- [Tech Stack](#tech-stack)
- [Basic Seeding Setup (Midcap Only)](#basic-seeding-setup-midcap-only)
- [Set Up Discord Bot (Midcap Only)](#set-up-discord-bot-optional)
- [Extended Seeding Setup (Last Two Lines Blocked)](#extended-seeding-setup-last-two-lines-blocked)
- [Set Up Discord Bot (Extended Version)](#set-up-discord-bot-extended-version)
- [Running Scripts Persistently with PM2](#running-scripts-persistently-with-pm2)
- [Usage](#usage)
- [Roadmap](#roadmap)
- [License](#license)
- [Contact](#contact)

> [!TIP]  
> To use both bots you need 2 bot directories and one main directory!
<div>
  <img src=https://github.com/FwSchultz/assets/blob/main/bots/2KU77B0N3S/geofence/geofence-struktur.png alt="logo" width="400" />
</div>

---

## About the Project

This repository (**HLL-GEOFENCE**) provides scripts for managing seeding configurations for Hell Let Loose (HLL) servers using geofencing. It includes two setups: Basic Seeding (Midcap only) and Extended Seeding (last two lines blocked). Both setups can be configured for different player counts and include Docker and Discord integration options. Below are the setup instructions for each.

---

## Tech Stack

- **Programming language:** JavaScript  
- **Libraries:** `discord.js`, `dotenv`, `node-fetch`, `requests`  
- **Dev-Libraries:** `nodemon`

---

## Basic Seeding Setup (Midcap Only)

1. Clone the Repository:

```bash
git clone https://github.com/2KU77B0N3S/hll-geofences
```

2. Rename the Folder:

```bash
mv hll-geofences hll-geofences-basic
cd hll-geofences-basic
```

3. Select Player Count Configuration:

```bash
# 40 Player
mv seeding.midcap.40player.config.yml config.yml
```
or
```bash
# 50 Player
mv seeding.midcap.50player.config.yml config.yml
```
or
```bash
# 60 Player
mv seeding.midcap.60player.config.yml config.yml
```

4. Edit Configuration:  
Open `config.yml` and fill in `SERVER-IP`, `RCON-PORT`, and `RCON-PW`.

5. Rename the Docker configuration file:

```bash
mv midcap.docker-compose.yml docker-compose.yml
```

6. Build Docker image:

```bash
docker compose build
```

7. Start Docker:

```bash
docker compose up -d
```

8. Stop Docker (optional):

```bash
docker compose down
```

---

## Set Up Discord Bot (Optional)

1. Rename and fill out the environment file (`.env`):

```bash
mv midcap.example.env .env
```

2. Rename the script:

```bash
mv midcap.main.mjs main.mjs
```

3. Install dependencies:

```bash
npm install
```

4. Start the script:

```bash
node main.mjs
```

---

> [!NOTE]  
> If you also want to use the extended version and are still in `hll-geofences-basic`, go back to the main directory:

```bash
cd ..
```

---

## Extended Seeding Setup (Last Two Lines Blocked)

1. Clone the Repository:

```bash
git clone https://github.com/2KU77B0N3S/hll-geofences
```

2. Rename the Folder:

```bash
mv hll-geofences hll-geofences-extended
cd hll-geofences-extended
```

3. Select Player Count Configuration:

```bash
# 60 Player
mv seeding.3caps.60player.config.yml config.yml
```
or
```bash
# 70 Player
mv seeding.3caps.70player.config.yml config.yml
```
or
```bash
# 80 Player
mv seeding.3caps.80player.config.yml config.yml
```

4. Edit Configuration:  
Open `config.yml` and fill in `SERVER-IP`, `RCON-PORT`, and `RCON-PW`.

5. Rename Docker config:

```bash
mv extended.docker-compose.yml docker-compose.yml
```

6. Build Docker image:

```bash
docker compose build
```

7. Start Docker:

```bash
docker compose up -d
```

8. Stop Docker:

```bash
docker compose down
```

---

## Set Up Discord Bot (Extended Version)

1. Rename and fill out the environment file:

```bash
mv extended.example.env .env
```

2. Rename the script:

```bash
mv extended.main.mjs main.mjs
```

3. Install dependencies:

```bash
npm install
```

4. Start the script:

```bash
node main.mjs
```

---

## Running Scripts Persistently with PM2

1. Install PM2 globally:

```bash
npm install -g pm2
```

2. Setup PM2 autostart:

```bash
pm2 startup
```

3. Start the scripts:

```bash
cd hll-geofences-basic
pm2 start main.mjs --name hll-geofence-basic

cd ../hll-geofences-extended
pm2 start main.mjs --name hll-geofence-extended
```

4. Save PM2 configuration:

```bash
pm2 save
```

5. (Optional) Start PM2 again on reboot:

```bash
pm2 startup
```

### Managing PM2 Processes

```bash
pm2 status       # View processes
pm2 monit        # Monitor live
pm2 stop id      # Stop by ID
pm2 restart id   # Restart by ID
pm2 delete id    # Delete by ID
pm2 resurrect    # Recover after reboot
```

---

## Usage

The bot runs as a Discord bot and can be controlled via buttons.

---

## Roadmap

* [ ] Language file support (`en`, `es`, `fr`, `de`)
* [ ] More seeding logic presets
* [x] Admin control panel

---

## License

This project is licensed under the **MIT License**.

---

## Contact

Created by **2KU77B0N3S**. For questions or suggestions, please create an issue on GitHub.

📧 **Contact:** [Discord](https://discord.com/users/785886867073400903)

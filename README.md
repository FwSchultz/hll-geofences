# HLL Geofences

This repository provides scripts for managing seeding configurations for Hell Let Loose (HLL) servers using geofencing. It includes two setups: **Basic Seeding** (Midcap only) and **Extended Seeding** (last two lines blocked). Both setups can be configured for different player counts and include Docker and Discord integration options. Below are the setup instructions for each.

## Prerequisites

- Git
- Docker (for Docker-based deployment)
- Node.js and npm (for Discord bot control)
- PM2 (for persistent script execution)
- Access to HLL server RCON (IP, port, password)

## Basic Seeding Setup (Midcap Only)

This setup configures the server for midcap-only gameplay with player counts of 40, 50, or 60.

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/2KU77B0N3S/hll-geofences
   ```

2. **Rename the Folder**:
   ```bash
   mv hll-geofences hll-geofences-basic
   cd hll-geofences-basic
   ```

3. **Select Player Count Configuration**:
   Choose one of the following commands based on the desired player count:
   ```bash
   mv seeding.midcap.40player.config.yml config.yml
   ```
   or
   ```bash
   mv seeding.midcap.50player.config.yml config.yml
   ```
   or
   ```bash
   mv seeding.midcap.60player.config.yml config.yml
   ```

4. **Edit Configuration**:
   Open `config.yml` and fill in `SERVER-IP`, `RCON-PORT`, and `RCON-PW`.

5. **Set Up Docker**:
   - Rename the Docker configuration file:
     ```bash
     mv midcap.docker-compose.yml docker-compose.yml
     ```
   - Build the Docker image (required after changing files):
     ```bash
     docker compose build
     ```
   - Start the Docker container:
     ```bash
     docker compose up -d
     ```
   - Stop the Docker container (when needed):
     ```bash
     docker compose down
     ```

6. **Set Up Discord Bot (Optional)**:
   - Rename the environment file:
     ```bash
     mv midcap.example.env .env
     ```
   - Rename the script:
     ```bash
     mv midcap.main.mjs main.mjs
     ```
   - Install dependencies:
     ```bash
     npm install
     ```
   - Start the script:
     ```bash
     node main.mjs
     ```

## Extended Seeding Setup (Last Two Lines Blocked)

This setup blocks the last two lines for gameplay with player counts of 60, 70, or 80.

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/2KU77B0N3S/hll-geofences
   ```

2. **Rename the Folder**:
   ```bash
   mv hll-geofences hll-geofences-extended
   cd hll-geofences-extended
   ```

3. **Select Player Count Configuration**:
   Choose one of the following commands based on the desired player count:
   ```bash
   mv seeding.3caps.60player.config.yml config.yml
   ```
   or
   ```bash
   mv seeding.3caps.70player.config.yml config.yml
   ```
   or
   ```bash
   mv seeding.3caps.80player.config.yml config.yml
   ```

4. **Edit Configuration**:
   Open `config.yml` and fill in `SERVER-IP`, `RCON-PORT`, and `RCON-PW`.

5. **Set Up Docker**:
   - Rename the Docker configuration file:
     ```bash
     mv extended.docker-compose.yml docker-compose.yml
     ```
   - Build the Docker image (required after changing files):
     ```bash
     docker compose build
     ```
   - Start the Docker container:
     ```bash
     docker compose up -d
     ```
   - Stop the Docker container (when needed):
     ```bash
     docker compose down
     ```

6. **Set Up Discord Bot (Optional)**:
   - Rename the environment file:
     ```bash
     mv extended.example.env .env
     ```
   - Rename the script:
     ```bash
     mv extended.main.mjs main.mjs
     ```
   - Install dependencies:
     ```bash
     npm install
     ```
   - Start the script:
     ```bash
     node main.mjs
     ```

## Running Scripts Persistently with PM2

To run either or both scripts in the background with automatic restarts, use PM2.

1. **Install PM2 Globally**:
   ```bash
   npm install -g pm2
   ```

2. **Set Up PM2 Autostart**:
   ```bash
   pm2 startup
   ```

3. **Start Scripts**:
   - For Basic Seeding:
     ```bash
     cd hll-geofences-basic
     pm2 start main.mjs --name hll-geofence-basic
     ```
   - For Extended Seeding:
     ```bash
     cd hll-geofences-extended
     pm2 start main.mjs --name hll-geofence-extended
     ```

4. **Save PM2 Configuration**:
   ```bash
   pm2 save
   ```

5. **Enable PM2 Autostart**:
   ```bash
   pm2 startup
   ```

## Managing PM2 Processes

- **View Running Processes**:
  ```bash
  pm2 status
  ```
- **Monitor Processes**:
  ```bash
  pm2 monit
  ```
- **Control Processes** (replace `id` with the process ID from `pm2 status`):
  - Stop: `pm2 stop id`
  - Restart: `pm2 restart id`
  - Start: `pm2 start id`
  - Delete: `pm2 delete id`
- **Recover Processes** (if not showing):
  ```bash
  pm2 resurrect
  ```

## Notes

- Ensure all configuration files (`config.yml`, `.env`) are correctly filled out before starting.
- Run `docker compose build` after modifying any Docker-related files to ensure changes are applied.
- Docker and Discord bot setups are optional and can be skipped if not needed.
- PM2 is recommended for production to ensure scripts run continuously.

For issues or contributions, please open an issue or pull request on the [GitHub repository](https://github.com/2KU77B0N3S/hll-geofences).

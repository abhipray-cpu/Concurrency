# Use the latest version of Alpine Linux as the base image
FROM alpine:latest

# Install Chromium and dependencies
RUN apk --no-cache add \
    chromium \
    nss \
    freetype \
    freetype-dev \
    harfbuzz \
    ca-certificates \
    ttf-freefont \
    nodejs \
    npm \
    yarn

# Set environment variables for Chromium
# This tells Puppeteer / chromedp to skip the download process.
# Use `chromium-browser` as the executable for Chromium.
ENV CHROME_BIN=/usr/bin/chromium-browser \
    PUPPETEER_SKIP_CHROMIUM_DOWNLOAD=true \
    CHROMEDP_DISABLE_GPU=false

# Create a directory named /app in the container's filesystem
RUN mkdir /app

# Copy the CRAWLER binary from the build context into the /app directory inside the container
COPY CRAWLER /app

# Set the command to execute when the container starts. Here, it runs the CRAWLER binary.
CMD ["/app/CRAWLER"]
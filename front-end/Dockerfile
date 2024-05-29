# Build stage
FROM node:lts as build
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install
COPY . .
# Increase Node.js memory limit during build
ENV NODE_OPTIONS="--max_old_space_size=4096"
RUN npm run build:prod

# Production stage
FROM nginx:stable-alpine as production
COPY --from=build /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]

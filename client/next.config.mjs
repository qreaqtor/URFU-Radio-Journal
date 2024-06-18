/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  images: {
    domains: [
      "158.160.135.237",
      "host.docker.internal",
    ],
  },
};

export default nextConfig;

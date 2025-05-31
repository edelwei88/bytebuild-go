import type { NextConfig } from 'next';

const nextConfig: NextConfig = {
  async redirects() {
    return [
      {
        source: '/',
        destination: '/landing',
        permanent: true,
      },
      {
        source: '/app',
        destination: '/app/compile',
        permanent: true,
      },
    ];
  },
};

export default nextConfig;

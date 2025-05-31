import type { Metadata } from 'next';
import './globals.css';
import { ubuntuSans } from '@/lib/fonts';
import { Toaster } from '@/components/ui/sonner';

export const metadata: Metadata = {
  title: 'Bytebuild',
  description: 'Compile code',
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang='en'>
      <body className={`${ubuntuSans.className} bg-stone-300 antialiased`}>
        {children}
        <Toaster />
      </body>
    </html>
  );
}

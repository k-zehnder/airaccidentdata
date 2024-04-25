'use client';

import './globals.css';
import { ThemeProvider } from '@/components/theme-provider';
import { cn } from '../lib/utils';
import Head from 'next/head';
import { usePathname } from 'next/navigation';

const fontStyling = {
  fontFamily: '"Inter", sans-serif',
};

const metadata = {
  title: 'airaccidentdata.com',
  description: 'Latest plane crash accidents and incident news!',
  siteUrl: 'http://airaccidentdata.com',
  ogImage:
    'https://s3.us-west-1.amazonaws.com/s.airaccidentdata.com/og-image.png',
};

interface RootLayoutProps {
  children: React.ReactNode;
}

export default function RootLayout({ children }: RootLayoutProps) {
  const pathname = usePathname();

  return (
    <html lang="en">
      <Head>
        <title>{metadata.title}</title>
        <meta name="robots" content="follow, index" />
        <meta name="description" content={metadata.description} />
        <meta property="og:url" content={`${metadata.siteUrl}${pathname}`} />
        <meta property="og:type" content="website" />
        <meta property="og:site_name" content={metadata.title} />
        <meta property="og:description" content={metadata.description} />
        <meta property="og:title" content={metadata.title} />
        <meta
          property="og:image"
          content={`${metadata.siteUrl}${metadata.ogImage}`}
        />
        <meta name="twitter:card" content="summary_large_image" />
        <meta name="twitter:title" content={metadata.title} />
        <meta name="twitter:description" content={metadata.description} />
        <meta
          name="twitter:image"
          content={`${metadata.siteUrl}${metadata.ogImage}`}
        />
        <link rel="canonical" href={`${metadata.siteUrl}${pathname}`} />
        <link
          rel="icon"
          href={`${metadata.siteUrl}/favicon.ico`}
          type="image/x-icon"
          sizes="16x16"
        />
      </Head>
      <body
        className={cn('min-h-screen bg-background antialiased')}
        style={fontStyling}
      >
        <ThemeProvider attribute="class" defaultTheme="system" enableSystem>
          <div className="relative flex min-h-screen flex-col">
            <div className="flex w-full flex-1 flex-col overflow-hidden">
              {children}
            </div>
          </div>
        </ThemeProvider>
      </body>
    </html>
  );
}

import './globals.css';
import { ThemeProvider } from '@/components/ThemeProvider';
import { cn } from '../lib/utils';
import { GoogleAnalytics } from '@next/third-parties/google';

const fontStyling = {
  fontFamily: '"Inter", sans-serif',
};

interface Metadata {
  title: string;
  description: string;
  faviconUrl: string;
  ogImageUrl: string;
}

export const metadata: Metadata = {
  title: 'airaccidentdata.com',
  description: 'Latest plane crash accidents and incident news!',
  faviconUrl: 'https://s.airaccidentdata.com/favicon.ico',
  ogImageUrl: 'https://s.airaccidentdata.com/og-image.png',
};

interface RootLayoutProps {
  children: React.ReactNode;
}

export default function RootLayout({ children }: RootLayoutProps) {
  return (
    <html lang="en">
      <head>
        <title>{metadata.title}</title>
        <meta name="description" content={metadata.description} />
        <meta property="og:title" content={metadata.title} />
        <meta property="og:description" content={metadata.description} />
        <meta property="og:image" content={metadata.ogImageUrl} />
        <link
          rel="icon"
          href={metadata.faviconUrl}
          type="image/x-icon"
          sizes="16x16"
        />
      </head>
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
        <GoogleAnalytics gaId="G-D9DT897WDG" />
      </body>
    </html>
  );
}

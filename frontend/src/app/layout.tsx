import "./globals.css";
import { ThemeProvider } from "@/components/theme-provider";
import { cn } from "../lib/utils";

// Define the font styling using CSS-in-JS or a separate CSS class
const fontStyling = {
  fontFamily: '"Inter", sans-serif',
};

export const metadata = {
  title: "airaccidentdata.com",
  description: "Latest plane crash accidents and incident news.",
};

interface RootLayoutProps {
  children: React.ReactNode;
}

export default async function RootLayout({ children }: RootLayoutProps) {

  return (
      <html lang="en">
        <body
          className={cn("min-h-screen bg-background antialiased")}
          style={fontStyling}
        >
          <ThemeProvider attribute="class" defaultTheme="system" enableSystem>
            <div className="relative flex min-h-screen flex-col">
              <main className="flex w-full flex-1 flex-col overflow-hidden">
                {children}
              </main>
            </div>
          </ThemeProvider>
        </body>
      </html>
   );
}

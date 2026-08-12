import type { Metadata, Viewport } from "next";
import { IBM_Plex_Mono, Manrope } from "next/font/google";
import "./globals.css";

const sans = Manrope({
  subsets: ["latin"],
  variable: "--font-sans",
  display: "swap",
});

const mono = IBM_Plex_Mono({
  subsets: ["latin"],
  weight: ["400", "500", "600"],
  variable: "--font-mono",
  display: "swap",
});

export const metadata: Metadata = {
  metadataBase: new URL("https://aihostcheck.bond"),
  title: {
    default: "AIHostCheck — Host context AI should not guess",
    template: "%s · AIHostCheck",
  },
  description:
    "Open-source, privacy-aware cross-OS diagnostics that give GPT and AI coding agents evidence about a developer host before they propose commands.",
  applicationName: "AIHostCheck",
  keywords: [
    "AI coding agent",
    "developer environment",
    "host diagnostics",
    "cross-platform",
    "machine-readable",
    "open source",
  ],
  alternates: {
    canonical: "/",
  },
  openGraph: {
    type: "website",
    url: "/",
    siteName: "AIHostCheck",
    title: "Host context AI should not guess",
    description:
      "Evidence-backed, machine-readable diagnostics for AI development workflows.",
  },
  twitter: {
    card: "summary_large_image",
    title: "AIHostCheck",
    description:
      "Host context AI should not guess. Open-source cross-OS diagnostics for AI development workflows.",
  },
};

export const viewport: Viewport = {
  width: "device-width",
  initialScale: 1,
  themeColor: "#07110f",
  colorScheme: "dark",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className={`${sans.variable} ${mono.variable}`}>
      <body>{children}</body>
    </html>
  );
}

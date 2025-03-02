import type { Metadata } from "next";
import { Toaster } from "sonner";

export const metadata: Metadata = {
  title: "Nimbus-Drive",
  description: "A self-hosted replica of Google Drive, Store your files with complete privacy and control",
};

export default function DriveLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <>
      {children}
      <Toaster />
    </>
  );
}
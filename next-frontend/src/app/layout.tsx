import type React from "react"
import { Inter } from "next/font/google"
import "./globals.css"
import { Header } from "@/components/header"

const inter = Inter({ subsets: ["latin"] })

export const metadata = {
  title: "Payment Gateway",
  description: "A full payment gateway simulation",
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="pt-BR">
      <body className={`${inter.className} bg-[#1a2332] min-h-screen`}>
        <Header />
        <main className="container mx-auto py-4">{children}</main>
        <footer className="py-4 text-center text-sm text-gray-400">
          © All rights reserved to Full Cycle Company.
        </footer>
      </body>
    </html>
  )
}


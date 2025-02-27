'use client'

import { Cloud, ArrowRight } from "lucide-react"
import { Button } from "@/components/ui/button"
import Link from "next/link"
import { useState } from "react"

function GetStarted() {
  const [isHovering, setIsHovering] = useState(false);
  
  return (
    <div className="container mx-auto px-4 sm:px-6 lg:px-8">
      <main className='flex flex-col items-center min-h-screen py-16 md:justify-center'>
        <div className="relative mb-8 md:mb-12">
          <div className="absolute -inset-1.5 rounded-full bg-gradient-to-r from-primary/20 to-primary/40 blur-xl"></div>
          <div className="relative flex items-center justify-center rounded-full bg-background p-4 md:p-5 shadow-lg">
            <Cloud className="h-8 w-8 md:h-12 md:w-12 text-primary" />
          </div>
        </div>
        
        <h1 className='text-5xl md:text-7xl font-bold font-primary text-center mb-2 md:mb-3'>Nimbus</h1>
        
        <h2 className='text-3xl md:text-5xl lg:text-6xl font-bold font-primary py-4 md:py-6 text-center'>
          Your Files, Your Control
        </h2>
        
        <p className='text-lg md:text-xl text-center max-w-3xl px-4 md:px-0 py-2 md:py-4 font-primary font-normal text-muted-foreground'>
          A self-hosted replica of Google Drive, Store your files with complete privacy and control
        </p>
        
        <Button 
          className="mt-8 md:mt-10 rounded-full font-medium text-base md:text-lg px-6 md:px-8 py-2.5 md:py-3 transition-all flex items-center gap-2"
          size="lg"
          onMouseEnter={() => setIsHovering(true)}
          onMouseLeave={() => setIsHovering(false)}
        >
          <Link href="/drive" className="flex items-center gap-2"> 
            Go to Drive
            <ArrowRight 
              className={`h-5 w-5 transition-transform duration-300 ${isHovering ? 'translate-x-1' : ''}`} 
            />
          </Link>
        </Button>
      </main>
    </div>
  )
}

export default GetStarted
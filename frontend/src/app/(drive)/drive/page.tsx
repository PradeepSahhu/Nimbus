"use client"

import React from 'react'
import { Button } from "@/components/ui/button"
import { useRouter } from 'next/navigation'
import { toast } from "sonner"
import { LogOut } from "lucide-react"

function DrivePage() {
  const router = useRouter()
  
  const handleLogout = async () => {
    try {
      const response = await fetch('/api/auth/logout', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        }
      })
      
      if (response.ok) {
        toast.success("Logged out successfully")
        router.push('/auth/login')
        router.refresh()
      } else {
        toast.error("Failed to logout")
      }
    } catch (error) {
      console.error("Logout error:", error)
      toast.error("An error occurred while logging out")
    }
  }
  
  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex items-center justify-between mb-8">
        <h1 className="text-2xl font-bold">My Drive</h1>
        
        <Button 
          variant="outline" 
          onClick={handleLogout}
          className="flex items-center gap-2"
        >
          <LogOut className="h-4 w-4" />
          Logout
        </Button>
      </div>
      
      <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 gap-4">
        {/* Drive content will go here */}
        <div className="border rounded-md p-4 bg-gray-50">
          <p className="text-center text-gray-500">No files yet</p>
        </div>
      </div>
    </div>
  )
}

export default DrivePage
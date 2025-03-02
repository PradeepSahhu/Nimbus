"use client"

import React, { useState } from 'react'
import Link from 'next/link'
import { Cloud } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { useRouter } from 'next/navigation'
import { toast } from "sonner"

function LoginPage() {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [isLoading, setIsLoading] = useState(false)
  const router = useRouter()

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setIsLoading(true)
    
    try {
      const response = await fetch('/api/auth/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      })

      const data = await response.json()

      if (!response.ok) {
        if (data.error === 'invalid credentials') {
          toast.error("Login failed", {
            description: "Invalid email or password. Please try again.",
          })
        } else {
          toast.error("Login failed", {
            description: data.error || "Something went wrong",
          })
        }
        setIsLoading(false)
        return
      }
      
      toast.success("Login successful", {
        description: "Welcome back!",
      })
      
      router.push('/drive')
      
      router.refresh()
    } catch (error) {
      console.error('Login error:', error)
      toast.error("Login failed", {
        description: "An unexpected error occurred. Please try again.",
      })
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <div className="container mx-auto px-4 sm:px-6 lg:px-8">
      <main className="flex flex-col items-center min-h-screen py-16 md:py-28">
        {/* Logo */}
        <div className="relative mb-8">
          <div className="absolute -inset-1 rounded-full bg-gradient-to-r from-primary/20 to-primary/40 blur-xl"></div>
          <div className="relative flex items-center justify-center rounded-full bg-background p-4 shadow-lg">
            <Cloud className="h-8 w-8 text-primary" />
          </div>
        </div>
        
        {/* Heading */}
        <h1 className="text-3xl md:text-4xl font-bold font-primary text-center mb-8">
          Login to Nimbus
        </h1>
        
        {/* Form */}
        <div className="w-full max-w-md bg-gray-200 rounded-lg shadow-sm p-6 md:p-8 border">
          <form onSubmit={handleSubmit} className="space-y-6">
            <div className="space-y-2">
              <Label htmlFor="email">Email</Label>
              <Input
                id="email"
                type="email"
                placeholder="you@example.com"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                required
                className="w-full bg-white"
              />
            </div>
            
            <div className="space-y-2">
              <div className="flex items-center justify-between">
                <Label htmlFor="password">Password</Label>
                <Link 
                  href="/auth/forgot-password"
                  className="text-sm text-primary hover:underline"
                >
                  Forgot password?
                </Link>
              </div>
              <Input
                id="password"
                type="password"
                placeholder="••••••••"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                required
                className="w-full bg-white"
              />
            </div>
            
            <Button
              type="submit"
              className="w-full rounded-md"
              disabled={isLoading}
            >
              {isLoading ? "Signing in..." : "Sign in"}
            </Button>
          </form>
          
          <div className="mt-6 text-center text-sm">
            <p>
              Don&apos;t have an account?{" "}
              <Link 
                href="/auth/register" 
                className="text-primary font-medium hover:underline"
              >
                Register
              </Link>
            </p>
          </div>
        </div>
      </main>
    </div>
  )
}

export default LoginPage
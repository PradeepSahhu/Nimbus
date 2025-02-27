"use client"

import React, { useState } from 'react'
import Link from 'next/link'
import { Cloud } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { useRouter } from 'next/navigation'
import { toast } from "sonner"

function RegisterPage() {
  const [username, setUsername] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [isLoading, setIsLoading] = useState(false)
  const router = useRouter()

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setIsLoading(true)
    
    try {
      const response = await fetch('/api/auth/register', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, email, password }),
      })
  
      const data = await response.json()
  
      if (!response.ok) {
        if (data.message === 'email already exists') {
          toast.error("Email already registered", {
            description: "This email address is already in use. Please use a different email or sign in.",
          })
        } else {
          toast.error("Registration failed", {
            description: data.error || "Something went wrong",
          })
        }
        setIsLoading(false)
        return
      }
  
      toast.success("Account created!", {
        description: "Your account has been created successfully.",
      })
      
      router.push('/auth/login')
    } catch (error) {
      console.error('Registration error:', error)
      toast.error("Registration failed", {
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
          Create an Account
        </h1>
        
        {/* Form */}
        <div className="w-full max-w-md bg-gray-200 rounded-lg shadow-sm p-6 md:p-8 border">
          <form onSubmit={handleSubmit} className="space-y-6">
            <div className="space-y-2">
              <Label htmlFor="username">Username</Label>
              <Input
                id="username"
                type="text"
                placeholder="johndoe"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                required
                className="w-full bg-white"
              />
            </div>
            
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
              <Label htmlFor="password">Password</Label>
              <Input
                id="password"
                type="password"
                placeholder="••••••••"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                required
                className="w-full bg-white"
                minLength={6}
              />
            </div>
            
            <Button
              type="submit"
              className="w-full rounded-md"
              disabled={isLoading}
            >
              {isLoading ? "Creating account..." : "Create account"}
            </Button>
          </form>
          
          <div className="mt-6 text-center text-sm">
            <p>
              Already have an account?{" "}
              <Link 
                href="/auth/login" 
                className="text-primary font-medium hover:underline"
              >
                Sign in
              </Link>
            </p>
          </div>
        </div>
      </main>
    </div>
  )
}

export default RegisterPage
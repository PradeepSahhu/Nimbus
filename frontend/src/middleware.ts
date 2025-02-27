import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';

export function middleware(request: NextRequest) {
  const path = request.nextUrl.pathname;
  
  const isProtectedRoute = path.startsWith('/drive');
  
  const isAuthRoute = path.startsWith('/auth/login') || path.startsWith('/auth/register');
  
  const token = request.cookies.get('token')?.value;

  if (isProtectedRoute && !token) {
    const loginUrl = new URL('/auth/login', request.url);
    loginUrl.searchParams.set('callbackUrl', path);
    return NextResponse.redirect(loginUrl);
  }
  
  if (isAuthRoute && token) {
    return NextResponse.redirect(new URL('/drive', request.url));
  }
  
  return NextResponse.next();
}

export const config = {
  matcher: [
    '/drive/:path*',
    '/auth/login',
    '/auth/register'
  ]
};
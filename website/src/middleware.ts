import { NextRequest, NextResponse } from 'next/server';

export async function middleware(req: NextRequest) {
  const res = NextResponse.next();
  if (!req.cookies.has('Authorization'))
    return NextResponse.redirect(new URL('/login', req.url));

  const Authorization = req.cookies.get('Authorization');
  const authStatus = await fetch('http://localhost:3001/auth', {
    method: 'get',
    headers: {
      Cookie: `${Authorization?.name}=${Authorization?.value}`,
    },
  });

  if (!authStatus.ok) return NextResponse.redirect(new URL('/login', req.url));

  return res;
}

export const config = {
  matcher: ['/app', '/app/:path*'],
};

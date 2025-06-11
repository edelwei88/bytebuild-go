import { NextRequest, NextResponse } from 'next/server';
import { User } from '@/types/api/user';

export async function middleware(req: NextRequest) {
  const res = NextResponse.next();
  if (!req.cookies.has('Authorization'))
    return NextResponse.redirect(new URL('/login', req.url));

  const Authorization = req.cookies.get('Authorization');
  const data = await fetch('http://localhost:3001/user/me', {
    method: 'get',
    headers: {
      Cookie: `${Authorization?.name}=${Authorization?.value}`,
    },
  });

  if (!data.ok) return NextResponse.redirect(new URL('/login', req.url));

  const user: User = await data.json();
  if (
    req.url.includes('/manager') &&
    !['admin', 'manager'].includes(user.role.name)
  )
    return NextResponse.redirect(new URL('/app/compile', req.url));

  if (req.url.includes('/admin') && !['admin'].includes(user.role.name))
    return NextResponse.redirect(new URL('/app/compile', req.url));

  return res;
}

export const config = {
  matcher: ['/app', '/app/:path*'],
};

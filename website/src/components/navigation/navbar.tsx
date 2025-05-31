import { VT323font } from '@/lib/fonts';
import { NavbarProps } from '@/types/props/navbar';
import { AppWindow, Brackets, KeyRound } from 'lucide-react';
import Link from 'next/link';

export async function Navbar({
  navbar,
  authorized = false,
}: {
  navbar: NavbarProps;
  authorized?: boolean;
}) {
  return (
    <header className='flex h-30 w-full items-center justify-center'>
      <nav className='flex h-15 w-full items-center justify-between rounded-3xl bg-black px-5 text-white transition-all duration-500 ease-in-out hover:shadow-2xl'>
        <div>
          <Link
            href={'/'}
            className={`${VT323font.className} text-4xl subpixel-antialiased`}
          >
            <span className='flex items-center justify-center gap-2'>
              <Brackets size={30} />
              BYTEBUILD
            </span>
          </Link>
        </div>
        <div className='flex items-center justify-between gap-5 text-xl'>
          {!authorized &&
            navbar.items.map(item => (
              <div className='hidden md:block' key={item.link}>
                <Link
                  href={item.link}
                  className='underline decoration-black underline-offset-8 transition-all duration-200 ease-in-out hover:decoration-white'
                >
                  {item.name}
                </Link>
              </div>
            ))}
          {!authorized && (
            <div className='md:hidden'>
              <Link href={navbar.items[0].link}>
                <KeyRound />
              </Link>
            </div>
          )}
          {authorized && (
            <div className='hidden md:block'>
              <Link
                href='/app'
                className='underline decoration-black underline-offset-8 transition-all duration-200 ease-in-out hover:decoration-white'
              >
                Приложение
              </Link>
            </div>
          )}
          {authorized && (
            <div className='md:hidden'>
              <Link href={navbar.items[0].link}>
                <AppWindow />
              </Link>
            </div>
          )}
        </div>
      </nav>
    </header>
  );
}

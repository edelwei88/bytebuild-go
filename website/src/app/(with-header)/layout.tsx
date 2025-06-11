import { Navbar } from '@/components/navigation/navbar';
import { NavbarProps } from '@/types/props/navbar';
import { cookies } from 'next/headers';

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const items: NavbarProps = {
    items: [
      {
        name: 'Войти',
        link: '/login',
      },
      {
        name: 'Зарегистрироваться',
        link: '/register',
      },
    ],
  };

  const cookieStore = await cookies();
  const Authorization = cookieStore.get('Authorization');
  const authStatus = await fetch('http://localhost:3001/user/me', {
    method: 'get',
    headers: {
      Cookie: `${Authorization?.name}=${Authorization?.value}`,
    },
  });

  return (
    <div className='mx-auto max-w-3xl min-w-[500px] px-10'>
      <Navbar navbar={items} authorized={authStatus.ok} />
      {children}
    </div>
  );
}

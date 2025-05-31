import { AppSidebar } from '@/components/sidebar/app-sidebar';
import {
  SidebarInset,
  SidebarProvider,
  SidebarTrigger,
} from '@/components/ui/sidebar';
import { User } from '@/types/api/user';
import { cookies } from 'next/headers';

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const cookieStore = await cookies();
  const Authorization = cookieStore.get('Authorization');
  const res = await fetch('http://localhost:3001/auth', {
    method: 'get',
    headers: {
      Cookie: `${Authorization?.name}=${Authorization?.value}`,
    },
  });
  const user: User = await res.json();
  const defaultOpen = cookieStore.get('sidebar_state')?.value === 'true';

  return (
    <SidebarProvider defaultOpen={defaultOpen} className='dark'>
      <AppSidebar user={user} />
      <SidebarInset>
        <>
          <header className='flex h-14 shrink-0 items-center gap-2'>
            <div className='flex flex-1 items-center gap-2 px-3'>
              <SidebarTrigger className='bg-white' />
            </div>
          </header>
          {children}
        </>
      </SidebarInset>
    </SidebarProvider>
  );
}

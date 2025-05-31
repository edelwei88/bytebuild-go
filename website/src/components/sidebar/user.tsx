'use client';

import { Check, ChevronsUpDown, CircleX } from 'lucide-react';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '../ui/dropdown-menu';
import {
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  useSidebar,
} from '../ui/sidebar';
import { toast } from 'sonner';
import { useRouter } from 'next/navigation';
import { User as UserType } from '@/types/api/user';

export function User({ user }: { user: UserType }) {
  const { isMobile } = useSidebar();
  const r = useRouter();

  const roleName = {
    user: 'Пользователь',
    manager: 'Менеджер',
  }[user.role.name];

  return (
    <SidebarMenu>
      <SidebarMenuItem>
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <div>
              <SidebarMenuButton
                size='lg'
                className='data-[state=open]:bg-sidebar-accent data-[state=open]:text-sidebar-accent-foreground'
              >
                <div className='grid flex-1 text-left text-sm leading-tight'>
                  <span className='truncate font-medium'>{user.username}</span>
                  <span className='truncate text-xs'>{user.email}</span>
                </div>
                <ChevronsUpDown className='ml-auto size-4' />
              </SidebarMenuButton>
            </div>
          </DropdownMenuTrigger>
          <DropdownMenuContent
            className='w-(--radix-dropdown-menu-trigger-width) min-w-56 rounded-lg'
            side={isMobile ? 'bottom' : 'right'}
            align='end'
            sideOffset={4}
          >
            <DropdownMenuLabel className='p-0 font-normal'>
              <div className='flex items-center gap-2 px-1 py-1.5 text-left text-sm'>
                <div className='grid flex-1 text-left text-sm leading-tight'>
                  <span className='truncate text-lg font-medium'>
                    {user.username}
                  </span>
                  <span className='truncate'>{user.email}</span>
                  <span className='truncate text-xs'>{roleName}</span>
                </div>
              </div>
            </DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuItem
              onClick={() => {
                async function Logout() {
                  try {
                    const res = await fetch('http://localhost:3001/logout', {
                      credentials: 'include',
                      method: 'get',
                    });
                    if (!res.ok) {
                      toast('Неверные данные', {
                        description: 'Отказано в доступе',
                        icon: <CircleX />,
                      });
                      r.push('/');
                    }
                    if (res.ok) {
                      toast('Успешный выход из аккаунта', {
                        duration: 1000,
                        icon: <Check />,
                      });
                      setTimeout(() => {
                        r.push('/');
                      }, 1500);
                    }
                  } catch (e) {
                    toast('Неожиданная ошибка');
                    console.log(e);
                  }
                }
                Logout();
              }}
            >
              Выйти
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </SidebarMenuItem>
    </SidebarMenu>
  );
}

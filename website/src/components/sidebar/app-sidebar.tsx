'use client';

import { VT323font } from '@/lib/fonts';
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from '../ui/sidebar';
import { User } from '@/types/api/user';
import { User as UserNode } from './user';
import { SidebarNav } from './sidebar-nav';
import Link from 'next/link';
import { Brackets } from 'lucide-react';

export function AppSidebar({
  user,
  ...props
}: React.ComponentProps<typeof Sidebar> & { user: User }) {
  const userMenuItems = [
    {
      name: 'Компилятор',
      href: '/app/compile',
    },
    {
      name: 'Мои компиляции',
      href: '/app/mycompiles',
    },
  ];
  const managerMenuItems = [
    {
      name: 'Список пользователей',
      href: '/app/manager/users',
    },
    {
      name: 'Список компиляторов',
      href: '/app/manager/compilers',
    },
    {
      name: 'Список компиляций',
      href: '/app/manager/compiles',
    },
  ];
  const adminMenuItems = [
    {
      name: 'Редактирование пользователей',
      href: '/app/admin/edit/users',
    },
  ];

  return (
    <Sidebar variant='inset' {...props}>
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton size='lg' asChild>
              <div className='cursor-pointer'>
                <Link
                  href='/'
                  className={`${VT323font.className} w-full text-3xl leading-tight subpixel-antialiased`}
                >
                  <span className='flex items-center gap-2'>
                    <Brackets size={30} />
                    BYTEBUILD
                  </span>
                </Link>
              </div>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>Приложение</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarNav items={userMenuItems} />
          </SidebarGroupContent>
        </SidebarGroup>
        {['manager', 'admin'].includes(user.role.name) && (
          <SidebarGroup>
            <SidebarGroupLabel>Администрирование</SidebarGroupLabel>
            <SidebarGroupContent>
              <SidebarNav items={managerMenuItems} />
            </SidebarGroupContent>
          </SidebarGroup>
        )}
        {user.role.name === 'admin' && (
          <SidebarGroup>
            <SidebarGroupLabel>Редактирование</SidebarGroupLabel>
            <SidebarGroupContent>
              <SidebarNav items={adminMenuItems} />
            </SidebarGroupContent>
          </SidebarGroup>
        )}
      </SidebarContent>
      <SidebarFooter>
        <UserNode user={user} />
      </SidebarFooter>
    </Sidebar>
  );
}

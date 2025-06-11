'use client';

import Link from 'next/link';
import { SidebarMenu, SidebarMenuButton, SidebarMenuItem } from '../ui/sidebar';
import { usePathname } from 'next/navigation';

export function SidebarNav({
  items,
}: {
  items: {
    name: string;
    href: string;
  }[];
}) {
  const p = usePathname();
  return (
    <SidebarMenu>
      {items.map(el => (
        <SidebarMenuItem key={el.name}>
          <SidebarMenuButton asChild isActive={p === el.href}>
            <Link href={el.href} className='text-sm'>
              {el.name}
            </Link>
          </SidebarMenuButton>
        </SidebarMenuItem>
      ))}
    </SidebarMenu>
  );
}

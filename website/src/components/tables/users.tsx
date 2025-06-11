'use client';

import { Compile } from '@/types/api/compile';
import { User } from '@/types/api/user';
import { ColumnDef } from '@tanstack/react-table';

export const usersColumns: ColumnDef<User>[] = [
  {
    accessorKey: 'id',
    header: 'ID',
  },
  {
    accessorKey: 'username',
    header: 'Имя пользователя',
  },
  {
    accessorKey: 'email',
    header: 'Email',
  },
  {
    accessorKey: 'role.name',
    header: 'Роль',
  },
  {
    accessorKey: 'compiles',
    header: 'Количество компиляций',
    cell: ({ row }) => {
      const compiles: Compile[] = row.getValue('compiles');
      return <div>{compiles.length}</div>;
    },
  },
];

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
    header: 'Username',
  },
  {
    accessorKey: 'email',
    header: 'Email',
  },
  {
    accessorKey: 'role.name',
    header: 'Role',
  },
  {
    accessorKey: 'compiles',
    header: 'Compiles',
    cell: ({ row }) => {
      const compiles: Compile[] = row.getValue('compiles');
      return <div>{compiles.length}</div>;
    },
  },
];

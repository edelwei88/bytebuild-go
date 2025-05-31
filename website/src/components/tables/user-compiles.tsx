'use client';

import { Compile } from '@/types/api/compile';
import { ColumnDef } from '@tanstack/react-table';

export const userCompilesColumns: ColumnDef<Compile>[] = [
  {
    accessorKey: 'id',
    header: 'ID',
  },
  {
    accessorKey: 'compiler.docker_image_name',
    header: 'Compiler',
  },
  {
    accessorKey: 'source_code',
    header: 'Source code',
    cell: ({ cell }) => (
      <div className='text-wrap'>{cell.renderValue() as string}</div>
    ),
  },
  {
    accessorKey: 'args',
    header: 'Args',
  },
  {
    accessorKey: 'exit_code',
    header: 'Exit code',
  },
  {
    accessorKey: 'stdout',
    header: 'Stdout',
    cell: ({ cell }) => (
      <div className='text-wrap'>{cell.renderValue() as string}</div>
    ),
  },
  {
    accessorKey: 'stderr',
    header: 'Stderr',
    cell: ({ cell }) => (
      <div className='text-wrap'>{cell.renderValue() as string}</div>
    ),
  },
];

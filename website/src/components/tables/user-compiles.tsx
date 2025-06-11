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
    header: 'Компилятор',
  },
  {
    accessorKey: 'source_code',
    header: 'Исходный код',
    cell: ({ cell }) => (
      <div className='text-wrap'>{cell.renderValue() as string}</div>
    ),
  },
  {
    accessorKey: 'args',
    header: 'Аргументы',
  },
  {
    accessorKey: 'exit_code',
    header: 'Код возврата',
  },
  {
    accessorKey: 'stdout',
    header: 'Стандартный вывод',
    cell: ({ cell }) => (
      <div className='text-wrap'>{cell.renderValue() as string}</div>
    ),
  },
  {
    accessorKey: 'stderr',
    header: 'Стандартный вывод ошибки',
    cell: ({ cell }) => (
      <div className='text-wrap'>{cell.renderValue() as string}</div>
    ),
  },
  {
    accessorKey: 'compile_time',
    header: 'Время компиляции',
    cell: ({ cell }) => {
      const dateTime = new Date((cell.getValue() as number) * 1000);
      const time = dateTime.toLocaleTimeString('ru-RU');
      const date = dateTime.toLocaleDateString('ru-RU');

      return <div className='text-wrap'>{`${time} ${date}`}</div>;
    },
  },
];

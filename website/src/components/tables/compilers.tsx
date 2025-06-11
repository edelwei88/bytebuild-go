import { ColumnDef } from '@tanstack/react-table';

export const compilersColumns: ColumnDef<{
  id: number;
  docker_image_name: string;
  language: string;
  file_extension: string;
}>[] = [
  {
    accessorKey: 'id',
    header: 'ID',
  },
  {
    accessorKey: 'docker_image_name',
    header: 'Название образа Docker',
  },
  {
    accessorKey: 'language',
    header: 'Язык',
  },
  {
    accessorKey: 'file_extension',
    header: 'Расширение файла',
  },
];

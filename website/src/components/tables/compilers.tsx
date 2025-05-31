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
    header: 'Docker image name',
  },
  {
    accessorKey: 'language',
    header: 'Language',
  },
  {
    accessorKey: 'file_extension',
    header: 'File extension',
  },
];

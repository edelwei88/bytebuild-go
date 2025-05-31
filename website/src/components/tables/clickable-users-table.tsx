'use client';
import { DataTable, DataTableProps } from './data-table';
import { useRouter } from 'next/navigation';

export function DataTableClickable<TData, TValue>({
  columns,
  data,
}: DataTableProps<TData, TValue>) {
  const r = useRouter();
  return (
    <DataTable
      columns={columns}
      data={data}
      onClick={rows => {
        r.push(`/app/users/${rows.getValue('id')}`);
      }}
    />
  );
}

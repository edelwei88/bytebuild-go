import { DataTable } from '@/components/tables/data-table';
import { userCompilesColumns } from '@/components/tables/user-compiles';
import { Compile } from '@/types/api/compile';
import { CircleArrowLeft } from 'lucide-react';
import { cookies } from 'next/headers';
import Link from 'next/link';

export default async function Page({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const result = await params;

  const cookiesStore = await cookies();
  const Authorization = cookiesStore.get('Authorization');
  const res = await fetch('http://localhost:3001/manager/users/compiles', {
    method: 'post',
    credentials: 'include',
    headers: {
      Cookie: `${Authorization?.name}=${Authorization?.value}`,
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      id: result.id,
    }),
  });
  const json: Compile[] = await res.json();

  return (
    <div className='mx-auto my-5 max-w-7/9 text-white'>
      <div className='mb-5 flex items-center justify-start gap-5'>
        <Link href='/app/manager/users'>
          <CircleArrowLeft />
        </Link>

        <span className='text-xl'>Компиляции пользовател с ID {result.id}</span>
      </div>
      <DataTable columns={userCompilesColumns} data={json} />
    </div>
  );
}

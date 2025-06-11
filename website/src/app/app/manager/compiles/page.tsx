import { compilesColumns } from '@/components/tables/compiles';
import { DataTable } from '@/components/tables/data-table';
import { Compile } from '@/types/api/compile';
import { cookies } from 'next/headers';

export default async function Page() {
  const cookiesStore = await cookies();
  const Authorization = cookiesStore.get('Authorization');
  const res = await fetch('http://localhost:3001/manager/compiles', {
    method: 'get',
    headers: {
      Cookie: `${Authorization?.name}=${Authorization?.value}`,
    },
  });
  const json: Compile[] = await res.json();

  return (
    <div className='mx-auto my-5 max-w-8/9 text-white'>
      <span className='text-xl'>Компиляции</span>
      <DataTable
        columns={compilesColumns}
        data={json.sort((a, b) => a.id - b.id)}
      />
    </div>
  );
}

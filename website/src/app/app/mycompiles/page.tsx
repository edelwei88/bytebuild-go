import { DataTableClickable } from '@/components/tables/clickable-user-table';
import { userCompilesColumns } from '@/components/tables/user-compiles';
import { User } from '@/types/api/user';
import { cookies } from 'next/headers';

export default async function Page() {
  const cookiesStore = await cookies();
  const Authorization = cookiesStore.get('Authorization');
  const res = await fetch('http://localhost:3001/auth', {
    method: 'get',
    headers: {
      Cookie: `${Authorization?.name}=${Authorization?.value}`,
    },
  });
  const json: User = await res.json();
  return (
    <div className='mx-auto my-5 max-w-7/9 text-white'>
      <DataTableClickable columns={userCompilesColumns} data={json.compiles} />
    </div>
  );
}

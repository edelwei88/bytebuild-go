import { DataTableClickable } from '@/components/tables/clickable-users-table';
import { usersColumns } from '@/components/tables/users';
import { User } from '@/types/api/user';
import { cookies } from 'next/headers';

export default async function Page() {
  const cookiesStore = await cookies();
  const Authorization = cookiesStore.get('Authorization');
  const res = await fetch('http://localhost:3001/manager/users', {
    method: 'get',
    headers: {
      Cookie: `${Authorization?.name}=${Authorization?.value}`,
    },
  });
  const json: User[] = await res.json();

  return (
    <div className='mx-auto my-5 max-w-7/9 text-white'>
      <span className='text-xl'>Пользователи</span>
      <DataTableClickable
        columns={usersColumns}
        data={json.sort((a, b) => a.id - b.id)}
      />
    </div>
  );
}

import { EditUsers } from '@/components/edit/edit-users';
import { Role } from '@/types/api/role';
import { User } from '@/types/api/user';
import { cookies } from 'next/headers';

export default async function Page() {
  const cookiesStore = await cookies();
  const Authorization = cookiesStore.get('Authorization');

  const usersRes = await fetch('http://localhost:3001/manager/users', {
    headers: {
      Cookie: `${Authorization?.name}=${Authorization?.value}`,
      'Content-Type': 'application/json',
    },
    credentials: 'include',
    method: 'get',
  });
  const usersJson: User[] = await usersRes.json();
  usersJson.sort((a, b) => a.id - b.id);

  const rolesRes = await fetch('http://localhost:3001/admin/roles', {
    headers: {
      Cookie: `${Authorization?.name}=${Authorization?.value}`,
      'Content-Type': 'application/json',
    },
    credentials: 'include',
    method: 'get',
  });
  const rolesJson: Role[] = await rolesRes.json();
  rolesJson.sort((a, b) => a.id - b.id);

  return (
    <div className='mx-auto min-w-3xl'>
      <EditUsers roles={rolesJson} users={usersJson} />
    </div>
  );
}

import { Compiler } from '@/components/compiler/compiler';
import { Language } from '@/types/api/language';
import { cookies } from 'next/headers';

export default async function Page() {
  const cookiesStore = await cookies();
  const Authorization = cookiesStore.get('Authorization');
  const res = await fetch('http://localhost:3001/user/langs', {
    method: 'get',
    headers: {
      Cookie: `${Authorization?.name}=${Authorization?.value}`,
    },
  });
  const json: Language[] = await res.json();

  return <Compiler languages={json} />;
}

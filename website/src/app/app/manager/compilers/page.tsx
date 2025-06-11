import { compilersColumns } from '@/components/tables/compilers';
import { DataTable } from '@/components/tables/data-table';
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
  const cs: {
    id: number;
    docker_image_name: string;
    language: string;
    file_extension: string;
  }[] = [];

  json.forEach(l => {
    l.compilers.forEach(c => {
      cs.push({
        id: c.id,
        docker_image_name: c.docker_image_name,
        language: l.name,
        file_extension: l.file_extension,
      });
    });
  });

  return (
    <div className='mx-auto my-5 max-w-7/9 text-white'>
      <span className='text-xl'>Компиляторы</span>
      <DataTable
        columns={compilersColumns}
        data={cs.sort((a, b) => a.id - b.id)}
      />
    </div>
  );
}

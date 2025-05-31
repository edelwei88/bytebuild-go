import { Compiler } from '@/components/compiler/compiler';
import { Language } from '@/types/api/language';
import { User } from '@/types/api/user';
import { cookies } from 'next/headers';

export default async function Page({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const result = await params;

  const resLangs = await fetch('http://localhost:3001/langs');
  const jsonLangs: Language[] = await resLangs.json();

  const cookiesStore = await cookies();
  const Authorization = cookiesStore.get('Authorization');
  const res = await fetch('http://localhost:3001/auth', {
    method: 'get',
    headers: {
      Cookie: `${Authorization?.name}=${Authorization?.value}`,
    },
  });
  const json: User = await res.json();

  const compile =
    json.compiles.find(c => c.id === parseInt(result.id)) ?? json.compiles[0];
  const lang =
    jsonLangs.find(l => l.id === compile.compiler.language_id) ?? jsonLangs[0];

  return (
    <Compiler
      languages={jsonLangs}
      languageRef={lang}
      argsRef={compile.args}
      compilerRef={compile.compiler}
      stderrRef={compile.stderr}
      sourceCodeRef={compile.source_code}
      stdoutRef={compile.stdout}
    />
  );
}

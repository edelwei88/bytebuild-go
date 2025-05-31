import { Compiler } from '@/components/compiler/compiler';
import { Language } from '@/types/api/language';

export default async function Page() {
  const res = await fetch('http://localhost:3001/langs');
  const json: Language[] = await res.json();

  return <Compiler languages={json} />;
}

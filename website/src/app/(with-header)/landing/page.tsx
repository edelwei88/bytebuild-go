import { About } from '@/components/landing/about';
import { Fast } from '@/components/landing/fast';
import { Interface } from '@/components/landing/interface';
import { LanguageSupport } from '@/components/landing/language-support';

export default function Page() {
  return (
    <div className='flex w-full flex-col gap-5'>
      <About />
      <LanguageSupport />
      <Fast />
      <Interface />
    </div>
  );
}

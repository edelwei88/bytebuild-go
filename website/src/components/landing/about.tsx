import Link from 'next/link';
import { Button } from '../ui/button';

export function About() {
  return (
    <div className='max-h-80 min-h-60 w-full rounded-3xl bg-black p-5'>
      <div className='flex min-h-60 flex-col items-center justify-between'>
        <div>
          <h1 className='max-w-fit scroll-m-20 text-4xl font-extrabold tracking-tight text-balance text-white'>
            Онлайн-компилятор программ
          </h1>
          <div className='mt-3'>
            <h2 className='w-full scroll-m-20 text-2xl text-white'>
              Пишите, компилируйте и тестируйте код в реальном времени!
            </h2>
          </div>
        </div>
        <Button
          className='cursor-pointer bg-white p-6 text-4xl text-black hover:bg-gray-300'
          asChild
        >
          <Link href='/app'>Попробовать сейчас</Link>
        </Button>
      </div>
    </div>
  );
}

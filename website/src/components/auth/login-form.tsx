/* eslint-disable react/no-children-prop */
'use client';
import { Button } from '@/components/ui/button';
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { useForm } from '@tanstack/react-form';
import Link from 'next/link';
import { z } from 'zod/v4';
import { FieldInfo } from './field-info';
import { toast } from 'sonner';
import { Check, CircleX } from 'lucide-react';
import { useRouter } from 'next/navigation';

export function LoginForm() {
  const r = useRouter();
  const formSchema = z.object({
    email: z.email('Введите правильную почту'),
    password: z.string().min(8, 'Минимальная длина пароля 8 символов'),
  });

  const form = useForm({
    defaultValues: {
      email: '',
      password: '',
    },
    validators: {
      onSubmit: formSchema,
    },
    onSubmit: ({ value }) => {
      async function Login() {
        try {
          const res = await fetch('http://localhost:3001/login', {
            headers: {
              'Content-Type': 'application/json',
            },
            credentials: 'include',
            method: 'post',
            body: JSON.stringify({
              email: value.email,
              password: value.password,
            }),
          });
          if (!res.ok)
            toast('Неверные данные', {
              description: 'Вы ввели неверный логин или пароль',
              icon: <CircleX />,
            });
          if (res.ok) {
            toast('Успешная авторизация', {
              duration: 1000,
              icon: <Check />,
            });
            setTimeout(() => {
              r.push('/app');
            }, 1500);
          }
        } catch (e) {
          toast('Неожиданная ошибка');
          console.log(e);
        }
      }
      Login();
    },
  });

  return (
    <div className='flex flex-col gap-6'>
      <Card className='dark bg-black'>
        <CardHeader>
          <CardTitle className='text-3xl text-white'>Авторизация</CardTitle>
          <CardDescription className='text-xl text-white'>
            Введите данные вашей учетной записи, чтобы авторизоваться
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form
            onSubmit={e => {
              e.preventDefault();
              e.stopPropagation();
            }}
          >
            <div className='flex flex-col gap-6'>
              <form.Field
                name='email'
                children={field => (
                  <div className='grid gap-3'>
                    <Label htmlFor='email' className='text-lg'>
                      Почта
                    </Label>
                    <Input
                      id='email'
                      type='email'
                      placeholder='user@gmail.com'
                      required
                      className='text-lg'
                      value={field.state.value}
                      onChange={e => field.handleChange(e.target.value)}
                    />
                    <FieldInfo field={field} />
                  </div>
                )}
              />
              <form.Field
                name='password'
                children={field => (
                  <div className='grid gap-3'>
                    <Label htmlFor='password' className='text-lg'>
                      Пароль
                    </Label>
                    <Input
                      id='password'
                      type='password'
                      required
                      className='text-lg'
                      value={field.state.value}
                      onChange={e => field.handleChange(e.target.value)}
                    />
                    <FieldInfo field={field} />
                  </div>
                )}
              />
              <div className='flex flex-col gap-3'>
                <Button
                  onClick={e => {
                    form.handleSubmit();
                    e.preventDefault();
                    e.stopPropagation();
                  }}
                  type='submit'
                  className='w-full text-xl'
                >
                  Войти
                </Button>
              </div>
              <div className='mt-4 text-center text-lg'>
                Нет учетной записи?{' '}
                <Link href='/register' className='underline underline-offset-4'>
                  Зарегистрироваться
                </Link>
              </div>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}

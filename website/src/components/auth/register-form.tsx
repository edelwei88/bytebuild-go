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
import { z } from 'zod/v4';
import { FieldInfo } from './field-info';
import { toast } from 'sonner';
import { Check, CircleX } from 'lucide-react';
import { useRouter } from 'next/navigation';

export function RegisterForm() {
  const r = useRouter();
  const formSchema = z.object({
    username: z
      .string()
      .min(8, 'Минимальная длина имени пользователя 8 символов'),
    email: z.email('Введите правильную почту'),
    password: z.string().min(8, 'Минимальная длина пароля 8 символов'),
  });

  const form = useForm({
    defaultValues: {
      username: '',
      email: '',
      password: '',
    },
    validators: {
      onSubmit: formSchema,
    },
    onSubmit: ({ value }) => {
      async function Register() {
        try {
          const res = await fetch('http://localhost:3001/auth/register', {
            headers: {
              'Content-Type': 'application/json',
            },
            credentials: 'include',
            method: 'post',
            body: JSON.stringify({
              username: value.username,
              email: value.email,
              password: value.password,
            }),
          });
          if (!res.ok)
            toast('Почта занята', {
              description: 'Введенная электронная почта занята',
              icon: <CircleX />,
            });
          if (res.ok) {
            toast('Успешная регистрация', {
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
      Register();
    },
  });

  return (
    <div className='flex flex-col gap-6'>
      <Card className='dark bg-black'>
        <CardHeader>
          <CardTitle className='text-3xl text-white'>Регистрация</CardTitle>
          <CardDescription className='text-xl text-white'>
            Зарегистрируйтесь, чтобы начать пользоваться приложением
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
                name='username'
                children={field => (
                  <div className='grid gap-3'>
                    <Label htmlFor='username' className='text-lg'>
                      Имя пользователя
                    </Label>
                    <Input
                      id='username'
                      type='username'
                      placeholder='username'
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
                  Зарегистрироваться
                </Button>
              </div>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}

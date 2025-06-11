/* eslint-disable react/no-children-prop */
'use client';

import { Role } from '@/types/api/role';
import { User } from '@/types/api/user';
import { useForm } from '@tanstack/react-form';
import { Check, CircleX } from 'lucide-react';
import { toast } from 'sonner';
import z from 'zod/v4';
import { Button } from '../ui/button';
import { Label } from '../ui/label';
import { Input } from '../ui/input';
import { FieldInfo } from '../auth/field-info';
import { Card, CardContent, CardHeader, CardTitle } from '../ui/card';
import { Select, SelectItem, SelectTrigger, SelectValue } from '../ui/select';
import { SelectContent } from '../ui/select';
import { useRouter } from 'next/navigation';
import { useState } from 'react';

export function EditUsers({ users, roles }: { users: User[]; roles: Role[] }) {
  const r = useRouter();
  const [id, setID] = useState(users[0]?.id ?? 0);

  const roleName = {
    user: 'Пользователь',
    manager: 'Менеджер',
    admin: 'Администратор',
  };

  const formSchema = z.object({
    id: z.int(),
    username: z
      .string()
      .min(8, 'Минимальная длина имени пользователя 8 символов'),
    email: z.email('Введите правильную почту'),
    password: z.string().refine(val => val.length >= 8 || val === '', {
      error: 'Минимальная длина пароля 8 символов',
    }),
    role: z.string(),
  });

  const form = useForm({
    defaultValues: {
      id: users[0]?.id ?? 0,
      username: users[0]?.username ?? '',
      email: users[0]?.email ?? '',
      password: '',
      role: users[0]?.role.name ?? '',
    },
    validators: {
      onSubmit: formSchema,
    },
    onSubmit: ({ value }) => {
      async function patchUser() {
        try {
          const res = await fetch('http://localhost:3001/admin/users/patch', {
            headers: {
              'Content-Type': 'application/json',
            },
            credentials: 'include',
            method: 'PATCH',
            body: JSON.stringify({
              id: value.id,
              username: value.username,
              email: value.email,
              password: value.password,
              role_id: roles.find(r => r.name === value.role)?.id ?? 1,
            }),
          });
          if (!res.ok)
            toast('Ошибка данных', {
              description: 'Вы ввели неверные данные',
              icon: <CircleX />,
            });
          if (res.ok) {
            toast('Успешное изменение данных', {
              duration: 1000,
              icon: <Check />,
            });
            r.refresh();
          }
        } catch (e) {
          toast('Неожиданная ошибка');
          console.log(e);
        }
      }
      patchUser();
    },
  });

  return (
    <div className='flex flex-col gap-6'>
      <Card className='bg-white'>
        <CardHeader>
          <CardTitle className='text-3xl text-black'>
            Изменение пользователя
          </CardTitle>
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
                name='id'
                children={field => (
                  <div className='grid gap-3'>
                    <Label htmlFor='id' className='text-lg text-black'>
                      ID
                    </Label>
                    <Select
                      value={field.state.value.toString()}
                      required
                      onValueChange={e => {
                        field.handleChange(parseInt(e));
                        form.state.values.email =
                          users.find(u => u.id.toString() === e)?.email ?? '';
                        form.state.values.role =
                          users.find(u => u.id.toString() === e)?.role.name ??
                          '';
                        form.state.values.username =
                          users.find(u => u.id.toString() === e)?.username ??
                          '';
                        form.state.values.password = '';
                        setID(parseInt(e));
                      }}
                    >
                      <SelectTrigger className='bg-black! text-white!'>
                        <SelectValue />
                      </SelectTrigger>
                      <SelectContent className='bg-black! text-white!'>
                        {users.map(u => (
                          <SelectItem key={u.id} value={u.id.toString()}>
                            {u.id}
                          </SelectItem>
                        ))}
                      </SelectContent>
                    </Select>
                    <FieldInfo field={field} />
                  </div>
                )}
              />
              <form.Field
                name='username'
                children={field => (
                  <div className='grid gap-3'>
                    <Label htmlFor='username' className='text-lg text-black'>
                      Имя пользователя
                    </Label>
                    <Input
                      id='username'
                      type='text'
                      required
                      className='bg-black! text-lg'
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
                    <Label htmlFor='email' className='text-lg text-black'>
                      Email
                    </Label>
                    <Input
                      id='email'
                      type='email'
                      required
                      className='bg-black! text-lg'
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
                    <Label htmlFor='password' className='text-lg text-black'>
                      Пароль
                    </Label>
                    <Input
                      id='password'
                      type='text'
                      required
                      className='bg-black! text-lg'
                      value={field.state.value}
                      onChange={e => field.handleChange(e.target.value)}
                    />
                    <FieldInfo field={field} />
                  </div>
                )}
              />
              <form.Field
                name='role'
                children={field => (
                  <div className='grid gap-3'>
                    <Label htmlFor='role' className='text-lg text-black'>
                      Роль
                    </Label>
                    <Select
                      value={field.state.value}
                      required
                      onValueChange={e => {
                        field.handleChange(e);
                      }}
                    >
                      <SelectTrigger className='bg-black! text-white!'>
                        <SelectValue />
                      </SelectTrigger>
                      <SelectContent className='bg-black! text-white!'>
                        {roles.map(r => (
                          <SelectItem key={r.id} value={r.name}>
                            {roleName[r.name as keyof typeof roleName]}
                          </SelectItem>
                        ))}
                      </SelectContent>
                    </Select>
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
                  className='w-full cursor-pointer text-xl'
                >
                  Изменить
                </Button>
                <Button
                  onClick={e => {
                    async function DeleteUser() {
                      try {
                        const res = await fetch(
                          `http://localhost:3001/admin/users/delete/${id}`,
                          {
                            credentials: 'include',
                            method: 'DELETE',
                          },
                        );
                        if (!res.ok)
                          toast('Ошибка данных', {
                            description: 'Вы ввели неверные данные',
                            icon: <CircleX />,
                          });
                        if (res.ok) {
                          toast('Успешное изменение данных', {
                            duration: 1000,
                            icon: <Check />,
                          });
                          r.refresh();
                          form.state.values.id = users[0]?.id ?? 0;
                          form.state.values.email = users[0]?.email ?? '';
                          form.state.values.role = users[0]?.role.name ?? '';
                          form.state.values.username = users[0]?.username ?? '';
                          form.state.values.password = '';
                        }
                      } catch (e) {
                        toast('Неожиданная ошибка');
                        console.log(e);
                      }
                    }
                    DeleteUser();
                    e.preventDefault();
                    e.stopPropagation();
                  }}
                  className='w-full cursor-pointer bg-red-400 text-xl'
                >
                  Удалить
                </Button>
              </div>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}

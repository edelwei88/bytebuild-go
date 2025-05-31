/* eslint-disable @typescript-eslint/no-explicit-any */
'use client';

import { Compiler as CompilerType } from '@/types/api/compiler';
import { Language } from '@/types/api/language';
import { Editor } from '@monaco-editor/react';
import { useRef, useState } from 'react';
import { LanguageSelect } from './language-select';
import { CompilerSelect } from './compiler-select';
import { Input } from '../ui/input';
import { Label } from '../ui/label';
import { Textarea } from '../ui/textarea';
import { Button } from '../ui/button';
import { toast } from 'sonner';
import { Check, CircleX } from 'lucide-react';

export function Compiler({
  languages,
  sourceCodeRef,
  stdoutRef,
  stderrRef,
  argsRef,
  compilerRef,
  languageRef,
}: {
  languages: Language[];
  sourceCodeRef?: string;
  stdoutRef?: string;
  stderrRef?: string;
  argsRef?: string;
  compilerRef?: CompilerType;
  languageRef?: Language;
}) {
  function handleEditorDidMount(editor: any) {
    editorRef.current = editor;
    editor.focus();
    editor.RenderMinimap = 0;
  }
  function handleOnChange(value: string | undefined) {
    if (value) setSourceCode(value);
  }

  const [sourceCode, setSourceCode] = useState(
    sourceCodeRef ?? 'console.log("Hello, world!")',
  );
  const [language, setLanguage] = useState<Language>(
    languageRef ?? languages[0],
  );
  const [compiler, setCompiler] = useState<CompilerType>(
    compilerRef ?? languages[0].compilers[0],
  );
  const [args, setArgs] = useState(argsRef ?? '');
  const [stdout, setStdout] = useState(stdoutRef ?? '');
  const [stderr, setStderr] = useState(stderrRef ?? '');
  const editorRef = useRef(null);

  return (
    <div className='flex h-full w-full flex-1 gap-5 p-5'>
      <div className='h-full w-1/2 overflow-clip rounded-2xl'>
        <Editor
          onMount={handleEditorDidMount}
          onChange={handleOnChange}
          defaultLanguage={language?.monaco_name}
          language={language?.monaco_name}
          defaultValue='console.log("Hello, world!")'
          value={sourceCode}
          options={{
            wordWrap: 'on',
            fontSize: 20,
            minimap: {
              enabled: false,
            },
            renderLineHighlight: 'none',
            scrollbar: {
              vertical: 'hidden',
            },
          }}
        />
      </div>
      <div className='flex h-full w-1/2 flex-col gap-5'>
        <div className='flex h-fit gap-5'>
          <Label className='text-xl text-white'>Язык</Label>
          <LanguageSelect
            langs={languages}
            language={language}
            setLanguage={setLanguage}
            setCompiler={setCompiler}
          />
          <Label className='text-xl text-white'>Компилятор</Label>
          <CompilerSelect
            compiler={compiler}
            lang={language}
            setCompiler={setCompiler}
          />
        </div>
        <div className='h-7/9'>
          <div>
            <Label className='mb-2 text-xl text-white'>Аргументы</Label>
            <Input
              type='text'
              className='bg-white! text-black!'
              value={args}
              onChange={e => setArgs(e.target.value)}
            />
          </div>
          <div className='flex h-full flex-col'>
            <Label className='mt-4 mb-2 text-xl text-white'>Stdout</Label>
            <Textarea
              className='h-full resize-none overflow-y-scroll bg-white! text-black! opacity-100!'
              disabled
              value={stdout}
            />
            <Label className='mt-4 mb-2 text-xl text-white'>Stderr</Label>
            <Textarea
              className='h-full resize-none overflow-y-scroll bg-white! text-black! opacity-100!'
              disabled
              value={stderr}
            />
          </div>
          <div>
            <Button
              className='mt-4 cursor-pointer text-xl'
              size='lg'
              onClick={() => {
                async function Compile() {
                  try {
                    const res = await fetch('http://localhost:3001/compile', {
                      headers: {
                        'Content-Type': 'application/json',
                      },
                      credentials: 'include',
                      method: 'post',
                      body: JSON.stringify({
                        language: language.name,
                        compiler: compiler.docker_image_name,
                        source_code: sourceCode,
                        args: args,
                      }),
                    });
                    if (!res.ok)
                      toast('Ошибка компиляции', {
                        description: 'Ошибка компиляции на сервере',
                        icon: <CircleX />,
                      });
                    if (res.ok) {
                      const json = await res.json();
                      setStdout(json.stdout);
                      setStderr(json.stderr);
                      toast('Успешная компиляция', {
                        duration: 1000,
                        icon: <Check />,
                      });
                    }
                  } catch (e) {
                    toast('Неожиданная ошибка');
                    console.log(e);
                  }
                }
                Compile();
              }}
            >
              Компилировать
            </Button>
          </div>
        </div>
      </div>
    </div>
  );
}

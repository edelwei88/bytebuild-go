'use client';

import { Compiler as CompilerType } from '@/types/api/compiler';
import { Language } from '@/types/api/language';
import { Editor } from '@monaco-editor/react';
import { LanguageSelect } from './language-select';
import { CompilerSelect } from './compiler-select';
import { Input } from '../ui/input';
import { Label } from '../ui/label';
import { Textarea } from '../ui/textarea';
import { Button } from '../ui/button';
import { Loader } from 'lucide-react';
import { useCompiler } from '@/hooks/use-compiler';

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
  const {
    sourceCode,
    language,
    setLanguage,
    compiler,
    setCompiler,
    args,
    setArgs,
    stdout,
    stderr,
    isCompiling,
    handleEditorDidMount,
    handleOnChange,
    compile,
  } = useCompiler({
    languages,
    sourceCodeRef,
    stdoutRef,
    stderrRef,
    argsRef,
    compilerRef,
    languageRef,
  });

  return (
    <div className='flex h-full w-full flex-1 flex-col gap-5 p-5 md:flex-row'>
      <div className='mb-5 h-80 w-full overflow-clip rounded-2xl md:mb-0 md:h-full md:w-1/2'>
        <Editor
          onMount={handleEditorDidMount}
          onChange={handleOnChange}
          defaultLanguage={language?.monaco_name}
          language={language?.monaco_name}
          defaultValue='console.log("Hello, world!")'
          value={sourceCode}
          options={{
            wordWrap: 'on',
            fontSize: 16,
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
      <div className='flex h-full w-full flex-col gap-5 md:w-1/2'>
        <div className='flex h-fit flex-wrap gap-3 md:gap-5'>
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
        <div className='flex flex-1 flex-col'>
          <div>
            <Label className='mb-2 text-xl text-white'>Аргументы</Label>
            <Input
              type='text'
              className='bg-white! text-black!'
              value={args}
              onChange={e => setArgs(e.target.value)}
            />
          </div>
          <div className='flex flex-1 flex-col'>
            <Label className='mt-4 mb-2 text-xl text-white'>
              Стандартный вывод
            </Label>
            <Textarea
              className='h-24 cursor-text! resize-none overflow-y-scroll bg-white! text-black! opacity-100! md:h-full'
              disabled
              value={stdout}
            />
            <Label className='mt-4 mb-2 text-xl text-white'>
              Стандартный вывод ошибки
            </Label>
            <Textarea
              className='h-24 cursor-text! resize-none overflow-y-scroll bg-white! text-black! opacity-100! md:h-full'
              disabled
              value={stderr}
            />
          </div>
          <div>
            <Button
              className='mt-4 w-full cursor-pointer text-xl md:w-auto'
              size='lg'
              onClick={compile}
            >
              {isCompiling ? (
                <span className='flex items-center justify-center gap-5'>
                  Компиляция <Loader />
                </span>
              ) : (
                'Компилировать'
              )}
            </Button>
          </div>
        </div>
      </div>
    </div>
  );
}

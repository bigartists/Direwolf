import {
  type ChangeEvent,
  useState,
  type CompositionEvent,
  useEffect,
} from 'react'

export default function useChineseInput(
  value: string,
  onChange: (value: string) => void
) {
  const [curValue, setValue] = useState(value)
  const [isComposing, setComposing] = useState(false) // 是否正在输入法录入中
  // 是否是谷歌浏览器
  const isChrome = navigator.userAgent.includes('WebKit')

  const handleCompositionStart = () => {
    setComposing(true)
  }

  const handleCompositionEnd = (
    event: CompositionEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    setComposing(false)
    // 谷歌浏览器onChange事件在handleCompositionEnd之前触发
    if (isChrome) {
      onChange?.(event.currentTarget.value)
    }
  }

  const handleChange = (
    event: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    // 没有采用输入法的值
    const rawValue = event.target.value
    setValue(rawValue)
    if (!isComposing) {
      onChange?.(rawValue)
    }
  }

  useEffect(() => {
    setValue(value)
  }, [value])

  return {
    curValue,
    handleChange,
    handleCompositionStart,
    handleCompositionEnd,
  }
}

import { useRef, useState, useLayoutEffect } from 'react';

export default function useScrollToBottom() {
  // for auto-scroll
  const scrollRef = useRef<HTMLDivElement>(null);
  const [autoScroll, setAutoScroll] = useState(true);
  const scrollToBottom = () => {
    const dom = scrollRef.current;
    if (dom) {
      setTimeout(() => {
        dom.scrollTop = dom.scrollHeight;
      }, 100);
    }
  };

  // auto scroll
  useLayoutEffect(() => {
    if (autoScroll) {
      scrollToBottom();
    }
  }, [autoScroll]);

  return {
    scrollRef,
    autoScroll,
    setAutoScroll,
    scrollToBottom,
  };
}

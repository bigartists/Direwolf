import { useRef, useState, useEffect, useCallback, useLayoutEffect } from 'react';

export default function useScrollToBottom(chats?: any[]) {
  // for auto-scroll
  const scrollRef = useRef<HTMLDivElement>(null);
  const [autoScroll, setAutoScroll] = useState(true);

  const scrollToBottom = () => {
    const dom = scrollRef.current;
    if (dom) {
      // setTimeout(() => (dom.scrollTop = dom.scrollHeight), 100)
      requestAnimationFrame(() => {
        dom.scrollTop = dom.scrollHeight;
      });
    }
  };

  // auto scroll
  useLayoutEffect(() => {
    autoScroll && scrollToBottom();
  });

  useEffect(() => {
    const dom = scrollRef.current;
    const handleScroll = () => {
      // 当用户往上滑动页至少30px时，禁用自动滚动
      if (dom && dom.scrollTop < dom.scrollHeight - dom.clientHeight - 30) {
        console.log('禁用自动滚动==================================');
        setAutoScroll(false);
      } else {
        setAutoScroll(true);
      }
    };
    if (chats && chats.length > 0) {
      dom?.addEventListener('scroll', handleScroll);
    }
    return () => {
      dom?.removeEventListener('scroll', handleScroll);
    };
  }, [chats]);

  return {
    scrollRef,
    autoScroll,
    setAutoScroll,
    scrollToBottom,
  };
}

// export default function useScrollToBottom(chats?: any[]) {
//   // for auto-scroll
//   const scrollRef = useRef<HTMLDivElement>(null)
//   const [autoScroll, setAutoScroll] = useState(true)
//   const scrollToBottom = () => {
//     const dom = scrollRef.current
//     if (dom) {
//       setTimeout(() => (dom.scrollTop = dom.scrollHeight), 100)
//     }
//   }

//   // auto scroll
//   useLayoutEffect(() => {
//     autoScroll && scrollToBottom()
//   })

//   return {
//     scrollRef,
//     autoScroll,
//     setAutoScroll,
//     scrollToBottom,
//   }
// }

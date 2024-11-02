---
nav:
  title: 组件
  path: /components
  order: 0
---

# CodeBlock

> 代码块组件，支持折叠、一键复制

```tsx
import { Markdown } from '@wair/components';

const text = '```cpp\n#include <iostream>\n\nint factorial(int N) {\n    if (N == 0) {\n        return 1;\n    } else {\n        int result = N;\n        while (N > 1) {\n            result *= --N;\n        }\n        return result;\n    }\n}\n\nint main() {\n    int N;\n    std::cout <<Enter an integer N:    std::cin >> N;\n    int fact = factorial(N);\n    std::cout <<The factorial of << N << is << fact << std::endl;\n    return 0;\n}\n```'

export default () => <Markdown content={text} />;
```


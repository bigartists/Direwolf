import type { CardProps } from '@mui/material/Card';

import Box from '@mui/material/Box';
import Card from '@mui/material/Card';

import { fPercent, fShortenNumber } from 'src/utils/format-number';

// ----------------------------------------------------------------------
import { Iconify } from 'src/components/iconify';

type Props = CardProps & {
  name?: string;
  title: string;
  total?: number;
  model?: string;
  percent: number;
  icon: React.ReactElement;
};

export function BookingWidgetSummary({
  name,
  title,
  percent,
  total,
  model,
  icon,
  sx,
  ...other
}: Props) {
  const renderTrending = (
    <Box gap={0.5} display="flex" alignItems="center" sx={{ typography: 'subtitle2' }}>
      <Iconify
        width={24}
        icon={
          percent < 0
            ? 'solar:double-alt-arrow-down-bold-duotone'
            : 'solar:double-alt-arrow-up-bold-duotone'
        }
        sx={{
          flexShrink: 0,
          color: 'success.main',
          ...(percent < 0 && { color: 'error.main' }),
        }}
      />
      <span>{model}</span>
    </Box>
  );

  return (
    <Card
      sx={{
        p: 2,
        pl: 3,
        display: 'flex',
        alignItems: 'center',
        ...sx,
      }}
      {...other}
    >
      <Box sx={{ flexGrow: 1 }}>
        <Box sx={{ color: 'text.secondary', typography: 'subtitle2' }}>{title}</Box>
        <Box sx={{ my: 1.5, typography: 'h4' }}>{name}</Box>

        {renderTrending}
      </Box>

      <Box
        sx={{
          width: 50,
          height: 50,
          lineHeight: 0,
          borderRadius: '50%',
          bgcolor: 'background.neutral',
        }}
      >
        {icon}
      </Box>
    </Card>
  );
}

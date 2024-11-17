import type { IconButtonProps } from '@mui/material/IconButton';

import { useState, useCallback } from 'react';

import Box from '@mui/material/Box';
import Stack from '@mui/material/Stack';
import Avatar from '@mui/material/Avatar';
import Drawer from '@mui/material/Drawer';
import Tooltip from '@mui/material/Tooltip';
import MenuItem from '@mui/material/MenuItem';
import { useTheme } from '@mui/material/styles';
import Typography from '@mui/material/Typography';
import IconButton from '@mui/material/IconButton';

import { paths } from 'src/routes/paths';
import { useRouter, usePathname } from 'src/routes/hooks';

import { varAlpha } from 'src/theme/styles';

import { Label } from 'src/components/label';
import { Iconify } from 'src/components/iconify';
import { Scrollbar } from 'src/components/scrollbar';
import { AnimateAvatar } from 'src/components/animate';

import { ICONS, ICONS_SRC } from '../config-nav-dashboard';
import { useGetConversations } from 'src/api/conversation';

// ----------------------------------------------------------------------

export type AccountDrawerProps = IconButtonProps & {
  handleClickSessionItem?: (session_id: string) => void;
};

export function ConversationDrawer({ handleClickSessionItem }: AccountDrawerProps) {
  const { conversations } = useGetConversations();

  const theme = useTheme();

  const router = useRouter();

  const pathname = usePathname();

  const [open, setOpen] = useState(false);

  const handleOpenDrawer = useCallback(() => {
    setOpen(true);
  }, []);

  const handleCloseDrawer = useCallback(() => {
    setOpen(false);
  }, []);

  const renderAvatar = (
    <AnimateAvatar
      width={96}
      slotProps={{
        avatar: { src: ICONS_SRC.banking },
        overlay: {
          border: 2,
          spacing: 3,
          color: `linear-gradient(135deg, ${varAlpha(theme.vars.palette.primary.mainChannel, 0)} 25%, ${theme.vars.palette.primary.main} 100%)`,
        },
      }}
    >
      Alex
    </AnimateAvatar>
  );

  const renderHead = (
    <Box display="flex" alignItems="center" sx={{ py: 2, pr: 1, pl: 2.5 }}>
      <Typography variant="inherit" sx={{ flexGrow: 1 }}>
        历史记录
      </Typography>

      <Tooltip title="Close">
        <IconButton onClick={handleCloseDrawer}>
          <Iconify icon="mingcute:close-line" />
        </IconButton>
      </Tooltip>
    </Box>
  );

  return (
    <>
      {/* <AccountButton
        onClick={handleOpenDrawer}
        photoURL={ICONS_SRC.chat}
        displayName={user?.displayName}
        sx={sx}
        {...other}
      /> */}

      <IconButton
        aria-label="settings"
        onClick={handleOpenDrawer}
        sx={{ p: 0, width: 40, height: 40 }}
      >
        {ICONS.history}
      </IconButton>

      <Drawer
        open={open}
        onClose={handleCloseDrawer}
        anchor="right"
        title="历史记录"
        slotProps={{ backdrop: { invisible: true } }}
        PaperProps={{ sx: { width: 320 } }}
      >
        {renderHead}

        <Scrollbar>
          <Stack
            sx={{
              py: 3,
              px: 2.5,
              borderTop: `dashed 1px ${theme.vars.palette.divider}`,
              // borderBottom: `dashed 1px ${theme.vars.palette.divider}`,
            }}
          >
            {conversations?.map((option) => {
              return (
                <MenuItem
                  key={option.session_id}
                  onClick={() =>
                    handleClickSessionItem && handleClickSessionItem(option?.session_id as string)
                  }
                  sx={{
                    py: 1,
                    color: 'text.secondary',
                    '& svg': { width: 24, height: 24 },
                    '&:hover': { color: 'text.primary' },
                  }}
                >
                  <Box component="span" sx={{ overflow: 'hidden' }}>
                    {option.title}
                  </Box>

                  {option.info && (
                    <Label color="error" sx={{ ml: 1 }}>
                      {option.info}
                    </Label>
                  )}
                </MenuItem>
              );
            })}
          </Stack>
        </Scrollbar>
      </Drawer>
    </>
  );
}

import Box from '@mui/material/Box';

import { DashboardContent } from 'src/layouts/dashboard';
import { BookingWidgetSummary } from './booking-widget-summary';
import { CustomBreadcrumbs } from 'src/components/custom-breadcrumbs';

// import { modelList } from 'src/.api-key';
import { isArray } from 'lodash';
import { IChatProps } from '../chat/type';
import { Avatar, Button } from '@mui/material';

import { Iconify } from 'src/components/iconify';
import { useBoolean } from 'src/hooks/use-boolean';
import { MaasNewForm } from './components/maas-new-form';
import { IModelItem } from 'src/types/common';
import { useGetMaasList } from 'src/api/maas';

// ----------------------------------------------------------------------

type Props = {
  title?: string;
};

export function ModelView({ title = 'Blank' }: Props) {
  const { maasList } = useGetMaasList();
  console.log('🚀 ~ MultiLLMChat ~ modelList:', maasList);

  const addModelForm = useBoolean();

  return (
    <DashboardContent maxWidth="xl">
      <CustomBreadcrumbs
        heading={title}
        links={[
          // { name: 'Dashboard', href: paths.dashboard.root },
          { name: 'basic ability' },
          { name: 'maas' },
        ]}
        action={
          <Button
            // component={RouterLink}
            // href={paths.dashboard.user.new}
            variant="contained"
            startIcon={<Iconify icon="mingcute:add-line" />}
            onClick={addModelForm.onTrue}
          >
            添加Maas
          </Button>
        }
        sx={{ mb: { xs: 3, md: 5 } }}
      />

      <Box
        sx={{
          gap: 3,
          display: 'grid',
          gridTemplateColumns: {
            xs: 'repeat(1, 1fr)',
            md: 'repeat(4, 1fr)',
          },
        }}
      >
        {isArray(maasList)
          ? maasList.map((model: IChatProps) => {
              return (
                <BookingWidgetSummary
                  key={model.model}
                  title="LLM"
                  percent={2.6}
                  name={model.name}
                  model={model.model}
                  sx={{
                    // bgcolor: 'background.neutral',
                    p: (theme) => theme.spacing(1.5, 1, 1.5, 2),
                  }}
                  icon={
                    <Avatar
                      src={model.avatar}
                      sx={{
                        width: 50,
                        height: 50,
                      }}
                    />
                  }
                />
              );
            })
          : null}
      </Box>
      <MaasNewForm open={addModelForm.value} onClose={addModelForm.onFalse} />
    </DashboardContent>
  );
}

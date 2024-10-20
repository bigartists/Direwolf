import type { IModelItem } from 'src/types/common';

import { z as zod } from 'zod';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { isValidPhoneNumber } from 'react-phone-number-input/input';

import Box from '@mui/material/Box';
import Stack from '@mui/material/Stack';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import LoadingButton from '@mui/lab/LoadingButton';
import DialogTitle from '@mui/material/DialogTitle';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';

import { Form, Field, schemaHelper } from 'src/components/hook-form';

// ----------------------------------------------------------------------

export type NewModelSchemaType = zod.infer<typeof NewModelSchema>;

export const NewModelSchema = zod.object({
  model_type: zod.string().min(1, { message: 'Model_type is required!' }),
  model: zod.string().min(1, { message: 'Model is required!' }),
  api_key: zod.string().min(1, { message: 'Api_key is required!' }),
  base_url: zod.string().min(1, { message: 'Base_url is required!' }),
});

// ----------------------------------------------------------------------

type Props = {
  open: boolean;
  onClose: () => void;
  onCreate: (address: IModelItem) => void;
};

export function ModelNewForm({ open, onClose, onCreate }: Props) {
  const defaultValues = {
    model: '',
    api_key: '',
    base_url: '',
    model_type: 'llm',
  };

  const methods = useForm<NewModelSchemaType>({
    mode: 'all',
    resolver: zodResolver(NewModelSchema),
    defaultValues,
  });

  const {
    handleSubmit,
    formState: { isSubmitting },
  } = methods;

  const onSubmit = handleSubmit(async (data) => {
    try {
      onCreate({
        model: data.model,
        model_type: data.model_type,
        api_key: data.api_key,
        base_url: data.base_url,
      });
      onClose();
    } catch (error) {
      console.error(error);
    }
  });

  return (
    <Dialog fullWidth maxWidth="sm" open={open} onClose={onClose}>
      <Form methods={methods} onSubmit={onSubmit}>
        <DialogTitle>添加模型</DialogTitle>

        <DialogContent dividers>
          <Stack spacing={3}>
            <Field.RadioGroup
              row
              label="Model Type"
              name="model_type"
              options={[
                { label: 'LLM', value: 'llm' },
                { label: 'Others', value: 'others' },
              ]}
            />
            <Field.Text name="model" label="Model" />

            <Field.Text name="api_key" label="Api Key" />

            <Field.Text name="base_url" label="Base Url" />
          </Stack>
        </DialogContent>

        <DialogActions>
          <Box sx={{ flexGrow: 1 }}>
            <Button variant="soft" color="error" onClick={onClose}>
              Delete
            </Button>
          </Box>

          <Button color="inherit" variant="outlined" onClick={onClose}>
            Cancel
          </Button>

          <LoadingButton type="submit" variant="contained" loading={isSubmitting}>
            Create
          </LoadingButton>
        </DialogActions>
      </Form>
    </Dialog>
  );
}

import React, { useState, useRef } from 'react';
import { Button } from '@/components/ui/button';
import { Upload, X } from 'lucide-react';
import { toast } from 'react-hot-toast';
import fileService from '@/services/fileService';
import { AxiosError } from 'axios';

interface FileUploadProps {
  onUploadComplete: () => void;
}

const FileUpload: React.FC<FileUploadProps> = ({ onUploadComplete }) => {
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [uploading, setUploading] = useState(false);
  const fileInputRef = useRef<HTMLInputElement>(null);

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files.length > 0) {
      setSelectedFile(e.target.files[0]);
    }
  };

  const handleUploadClick = () => {
    fileInputRef.current?.click();
  };

  const handleClearFile = () => {
    setSelectedFile(null);
    if (fileInputRef.current) {
      fileInputRef.current.value = '';
    }
  };

  const handleUpload = async () => {
    if (!selectedFile) {
      toast.error('Please select a file first');
      return;
    }

    setUploading(true);
    try {
      await fileService.uploadFile(selectedFile);
      toast.success('File uploaded successfully');
      setSelectedFile(null);
      onUploadComplete();
    } catch (error: unknown) {
      console.error('Upload error:', error);
      
      // Type guard to check if the error is an AxiosError
      if (error instanceof AxiosError) {
        const errorMessage = error.response?.data?.error || 'Failed to upload file';
        toast.error(errorMessage);
      } else {
        toast.error('An unexpected error occurred during upload');
      }
    } finally {
      setUploading(false);
    }
  };

  return (
    <div className="mb-6 p-4 border border-dashed rounded-md bg-gray-50">
      <input
        type="file"
        ref={fileInputRef}
        onChange={handleFileChange}
        className="hidden"
      />
      
      {selectedFile ? (
        <div className="flex flex-col gap-3">
          <div className="flex items-center justify-between">
            <div className="flex items-center gap-2">
              <div className="text-sm font-medium">{selectedFile.name}</div>
              <div className="text-xs text-gray-500">
                ({fileService.formatFileSize(selectedFile.size)})
              </div>
            </div>
            <Button
              variant="ghost"
              size="sm"
              onClick={handleClearFile}
              disabled={uploading}
            >
              <X className="h-4 w-4" />
            </Button>
          </div>
          <Button 
            onClick={handleUpload} 
            disabled={uploading}
            className="w-full"
          >
            {uploading ? 'Uploading...' : 'Upload File'}
          </Button>
        </div>
      ) : (
        <div className="flex flex-col items-center justify-center py-8">
          <Upload className="h-10 w-10 text-gray-400 mb-2" />
          <p className="text-sm text-gray-500 mb-2">Drag and drop a file or click to browse</p>
          <Button onClick={handleUploadClick}>Select File</Button>
        </div>
      )}
    </div>
  );
};

export default FileUpload;
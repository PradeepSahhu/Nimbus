"use client";

import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { Button } from '@/components/ui/button';
import { LogOut, RefreshCw } from 'lucide-react';
import { toast } from 'react-hot-toast';
import FileService, { FileItem } from '@/services/fileService';
import FileCard from '@/components/FileCard';
import FileUpload from '@/components/FileUpload';
import axios from 'axios';

export default function Drive() {
  const router = useRouter();
  const [files, setFiles] = useState<FileItem[]>([]);
  const [loading, setLoading] = useState(true);

  const fetchFiles = async () => {
    setLoading(true);
    try {
      const fetchedFiles = await FileService.listFiles();
      setFiles(fetchedFiles);
    } catch (error: any) {
      console.error('Error fetching files:', error);
      if (error?.response?.status === 401) {
        router.push('/auth/login');
        return;
      }
      toast.error('Failed to load files');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchFiles();
  }, []);

  const handleLogout = async () => {
    try {
      const response = await axios.post('/api/auth/logout');
      
      if (response.status === 200) {
        toast.success("Logged out successfully")
        router.push('/auth/login')
        router.refresh()
      } else {
        toast.error("Failed to logout")
      }
    } catch (error) {
      console.error("Logout error:", error)
      toast.error("An error occurred while logging out")
    }
  }

  const handleDeleteFile = async (fileId: string) => {
    try {
      await FileService.deleteFile(fileId);
      toast.success('File deleted successfully');
      fetchFiles();
    } catch (error) {
      console.error('Error deleting file:', error);
      toast.error('Failed to delete file');
    }
  };
  
  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex items-center justify-between mb-8">
        <h1 className="text-2xl font-bold">My Drive</h1>
        
        <div className="flex gap-2">
          <Button 
            variant="outline" 
            onClick={fetchFiles}
            disabled={loading}
            className="flex items-center gap-2"
          >
            <RefreshCw className={`h-4 w-4 ${loading ? 'animate-spin' : ''}`} />
            Refresh
          </Button>
          
          <Button 
            variant="outline" 
            onClick={handleLogout}
            className="flex items-center gap-2"
          >
            <LogOut className="h-4 w-4" />
            Logout
          </Button>
        </div>
      </div>
      
      <FileUpload onUploadComplete={fetchFiles} />
      
      {loading ? (
        <div className="text-center py-10">
          <RefreshCw className="h-8 w-8 animate-spin mx-auto text-gray-400" />
          <p className="mt-2 text-gray-500">Loading files...</p>
        </div>
      ) : files.length > 0 ? (
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
          {files.map(file => (
            <FileCard 
              key={file.id} 
              file={file} 
              onDelete={handleDeleteFile} 
            />
          ))}
        </div>
      ) : (
        <div className="border rounded-md p-12 bg-gray-50">
          <p className="text-center text-gray-500">No files yet. Upload your first file above.</p>
        </div>
      )}
    </div>
  );
}
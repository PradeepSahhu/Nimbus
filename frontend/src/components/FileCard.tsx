import React, { useState } from 'react';
import { Card, CardContent } from "@/components/ui/card";
import { Button } from '@/components/ui/button';
import { Download, Trash2, File, Image, FileText, Video, Music, Archive, Loader2 } from 'lucide-react';
import fileService, { FileItem } from '@/services/fileService';
import { formatDistanceToNow } from 'date-fns';
import { toast } from 'react-hot-toast';

interface FileCardProps {
  file: FileItem;
  onDelete: (id: string) => void;
}

const FileCard: React.FC<FileCardProps> = ({ file, onDelete }) => {
  const [isDownloading, setIsDownloading] = useState(false);
  const [isDeleting, setIsDeleting] = useState(false);
  
  const getFileIcon = () => {
    switch (file.type) {
      case 'image':
        return <Image className="h-8 w-8 text-blue-500" />;
      case 'document':
        return <FileText className="h-8 w-8 text-green-500" />;
      case 'video':
        return <Video className="h-8 w-8 text-red-500" />;
      case 'audio':
        return <Music className="h-8 w-8 text-purple-500" />;
      case 'archive':
        return <Archive className="h-8 w-8 text-yellow-500" />;
      default:
        return <File className="h-8 w-8 text-gray-500" />;
    }
  };

  const handleDownload = async () => {
    setIsDownloading(true);
    try {
      await fileService.downloadFile(file.id, file.name);
    } catch (error) {
      console.error('Download error:', error);
      toast.error('Failed to download file');
    } finally {
      setIsDownloading(false);
    }
  };

  const handleDelete = () => {
    setIsDeleting(true);
    onDelete(file.id);
  };
  
  const formatDate = (dateString: string) => {
    try {
      const date = new Date(dateString);
      
      if (isNaN(date.getTime())) {
        console.error('Invalid date:', dateString);
        return 'Unknown date';
      }
      
      return formatDistanceToNow(date, { addSuffix: true });
    } catch (error) {
      console.error('Date formatting error:', error, dateString);
      return 'Unknown date';
    }
  };

  return (
    <Card className="hover:shadow-md transition-shadow bg-gray-300">
      <CardContent className="p-4">
        <div className="flex items-center justify-center mb-4 pt-2">
          {getFileIcon()}
        </div>
        <h3 className="font-medium text-sm truncate mb-1" title={file.name}>
          {file.name}
        </h3>
        <div className="text-xs text-gray-500 mb-3">
          <p>{fileService.formatFileSize(file.size)}</p>
          <p>{formatDate(file.created_at)}</p>
        </div>
        <div className="flex justify-between gap-2 mt-2">
          <Button
            variant="outline"
            size="sm"
            className="flex-1 flex items-center justify-center"
            onClick={handleDownload}
            disabled={isDownloading}
          >
            {isDownloading ? (
              <Loader2 className="h-3 w-3 mr-1 animate-spin" />
            ) : (
              <Download className="h-3 w-3 mr-1" />
            )}
            <span>{isDownloading ? 'Downloading...' : 'Download'}</span>
          </Button>
          <Button
            variant="outline"
            size="sm"
            className="flex items-center justify-center text-red-500 hover:text-red-700"
            onClick={handleDelete}
            disabled={isDeleting}
          >
            {isDeleting ? (
              <Loader2 className="h-3 w-3 animate-spin" />
            ) : (
              <Trash2 className="h-3 w-3" />
            )}
          </Button>
        </div>
      </CardContent>
    </Card>
  );
};

export default FileCard;